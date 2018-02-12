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
)

type Users struct {

}

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

		result := TribonacciThroughCache(argument)


		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
}

func obtainArgument(path string) (int, error) {
	fragments := strings.Split(path,"/")
	lastFragment := fragments[len(fragments)-1]
	argument, err := strconv.Atoi(lastFragment)

	if argument < 0 {
		err = errors.New("negative arguments not supported")
	}

	//if argument > 100000 {
	//	err = errors.New("аргументы больше 100000 запрещены в демонстрационных целях")
	//}

	return argument, err
}

func TribonacciThroughCache(argument int ) string {
	result := TribonacсiIteroBig(argument)
	return  result.String()
}


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
		stepsToDone := argument-3 // опускаем уже просуммированные члены

		for i := 1; i < stepsToDone; i++ {

			// смещаемся по последовательности на один шаг вперёд
			first = second
			second = third
			third = next

			// суммируем все три члена
			sf := new(big.Int).Add(first,second)
			tsf := new(big.Int).Add(sf,third)

			next = tsf // запоминаем
		}

		return next
	}
}
