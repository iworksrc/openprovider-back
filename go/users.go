package openprovider

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"openprovider-back/go/models"
	"math/big"
)

type Users struct {

}

func GetTribonacсiValue(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")


		path := r.URL.Path

		argument, err := obtainArgument(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e := models.ErrorMessage{ Code: "400", Message: "Invalid argument"}
			errorMessage, _ := json.Marshal(e)
			http.Error(w, string(errorMessage) , http.StatusBadRequest)
			return
		}

		result := tribonacciThroughCache(argument)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
}

func obtainArgument(path string) (int, error) {
	fragments := strings.Split(path,"/")
	lastFragment := fragments[len(fragments)-1]
	return strconv.Atoi(lastFragment)
}

func tribonacciThroughCache(argument int ) string {

	//==== выбрать нужный метод вычисления ====
	//result := tribonacciRecursive(argument)
	result := tribonacсiItero(argument)

	//result := tribonacciRecursiveBig(argument)

	return  strconv.Itoa(result)
	//return  result.String()
}

func tribonacciRecursive(argument int) int {

	if argument == 0 {
		return 0
	}else if argument == 1{
		return 0
	} else if argument == 2{
		return 1
	} else if argument == 3{
		return 1
	} else{
		return  tribonacciRecursive(argument - 1) + tribonacciRecursive(argument - 2) + tribonacciRecursive(argument -3)
	}
}

func tribonacсiItero(argument int) int {
	var first = 0
	var second = 1
	var third = 1

	if argument == 0 {
		return 0
	}else if argument == 1 {
		return first
	}else if argument == 2 {
		return second
	}else if argument == 3 {
		return third
	}else {
		next := first + second + third

		stepsToDone := argument-3 // опускаем уже просуммированные члены
		for i := 1; i < stepsToDone; i++ {
			first = second
			second = third
			third = next
			next = first + second + third // вычисляем следующий член последовательности
		}
		return next
	}
}

func tribonacciRecursiveBig(argument int) *big.Int {

	if argument == 0 {
		return new(big.Int).SetUint64(0)
	}else if argument == 1{
		return new(big.Int).SetUint64(0)
	} else if argument == 2{
		return new(big.Int).SetUint64(0)
	} else if argument == 3{
		return new(big.Int).SetUint64(1)
	} else{
		f := tribonacciRecursiveBig(argument - 1)
		s := tribonacciRecursiveBig(argument - 2)
		t := tribonacciRecursiveBig(argument - 3)
		s.Add(s,t)
		f.Add(f,s)
		return f
	}
}
