package openprovider

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"openprovider-back/go/models"
	"math/big"
	"errors"
	"time"
	"github.com/patrickmn/go-cache"
)

// In-Memory cache with preset time-out limits
var memoryCache = cache.New(5*time.Minute, 10*time.Minute)

// Entrypoint /api/v1/openprovider/tribonachi/{argument}
func GetTribonacсiValue(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		path := r.URL.Path

		argument, err := obtainArgument(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e := models.ErrorMessage{ Code: "400", Message: "Invalid argument: " + err.Error()}
			errorMessage, _ := json.Marshal(e)
			http.Error(w, string(errorMessage) , http.StatusBadRequest)
			return
		}

		result := TribonacciThroughCache(argument, 100000)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
}


// Extract, validate and convert to a suitable
// form the argument from the query path. Detects negative
// values and values are more than a million (you can remove the check
// to the upper limit). In the case of non-valid transformation
// to int - it also returns an error
func obtainArgument(path string) (int, error) {
	fragments := strings.Split(path,"/")
	lastFragment := fragments[len(fragments)-1]
	argument, err := strconv.Atoi(lastFragment)

	if argument < 0 {
		err = errors.New("negative arguments not supported")
	}

	// can be commented out for stress tests
	if argument > 1000000 {
		err = errors.New("arguments greater than 1,000,000 are banned for demonstration purposes")
	}

	return argument, err
}


// Returns the value of the nth member of the Triboacci series.
// argument - is a member of the Triboacci series whose value is to be obtained.
// beginCacheLimit - value that specifies the initial limit of the argument
// starting with the calculated values will be cached
func TribonacciThroughCache(argument int, beginCacheLimit int ) string {

	// if the argument is large enough
	if argument > beginCacheLimit {
		// look in the cache
		value, found := memoryCache.Get(strconv.Itoa(argument))
		if found {
			return value.(string)
		}
	}

	// for small values of the argument or empty cache - calculate
	result := TribonacсiIteroBig(argument)

	// if the argument is large enough - remember
	if argument > beginCacheLimit {
		memoryCache.Set(strconv.Itoa(argument), result.String(), cache.DefaultExpiration)
	}

	return  result.String()
}

// Calculates the value of the nth member of the Triboacci series.
// An iterative calculation algorithm is used.
// To store the calculation results, use
// type of data big.Int is not a bounding dimension
// (the size of a variable of type big.Int is limited
// only the size of the machine's RAM).
// As a consequence, the function can calculate the values of the series
// from rather large values of the argument
// (one million or more) for acceptable time.
func TribonacсiIteroBig(argument int) *big.Int {
	var zero = new(big.Int).SetUint64(0)
	var first = new(big.Int).SetUint64(0)
	var second = new(big.Int).SetUint64(1)
	var third = new(big.Int).SetUint64(1)

	if argument == 0 {
		return zero
	}else if argument == 1 {
		return first
	}else if argument == 2 {
		return second
	}else if argument == 3 {
		return third
	}else {
		next := new(big.Int).SetUint64(2)
		stepsToDone := argument-3 // omit already summed terms

		for i := 1; i < stepsToDone; i++ {

			// move along the sequence one step forward
			first = second
			second = third
			third = next

			// summarize all three terms
			sf := new(big.Int).Add(first,second)
			tsf := new(big.Int).Add(sf,third)

			next = tsf // remember
		}

		return next
	}
}
