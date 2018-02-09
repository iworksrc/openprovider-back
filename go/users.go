package openprovider

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"openprovider-back/go/models"
	"math"
	"errors"
)

type Users struct {

}

var ErrorOverflow = errors.New("unit64 overflow")

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

		result, err := TribonacciThroughCache(argument)
		if err != nil {
			w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
			e := models.ErrorMessage{ Code: "416", Message: "Argument too big"}
			errorMessage, _ := json.Marshal(e)
			http.Error(w, string(errorMessage) , http.StatusRequestedRangeNotSatisfiable)
			return
		}


		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, result)
}

func obtainArgument(path string) (int, error) {
	fragments := strings.Split(path,"/")
	lastFragment := fragments[len(fragments)-1]
	return strconv.Atoi(lastFragment)
}

func TribonacciThroughCache(argument int ) (string, error) {

	//TODO some cache checking (не реализовано)

	result, err := tribonacciIteroOverflowProtected(argument)

	return  strconv.FormatUint(result, 10), err
}

func tribonacciIteroOverflowProtected(argument int) (uint64, error)  {
	var first uint64 = 0
	var second uint64 = 1
	var third uint64 = 1
	var tmp uint64

	if argument == 0 {
		return 0, nil
	}else if argument == 1 {
		return first, nil
	}else if argument == 2 {
		return second, nil
	}else if argument == 3 {
		return third, nil
	}else {
		next := first + second + third

		stepsToDone := argument-3 // опускаем уже просуммированные члены
		for i := 1; i < stepsToDone; i++ {
			first = second
			second = third
			third = next

			if second > (math.MaxUint64 - third) { // проверка на переполнение  second + third
				return 0, ErrorOverflow
			}
			tmp = second + third

			if first > (math.MaxUint64 - tmp) { // проверка на переполнение first + (second + third)
				return 0, ErrorOverflow
			}

			next = first + tmp // вычисляем следующий член последовательности
		}

		return next, nil
	}
}