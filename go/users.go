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

// Кеш распложенный в оперативной памяти с заданными границами сброса по времени
var memoryCache = cache.New(5*time.Minute, 10*time.Minute)

// Entrypoint.
// Обработка запроса к /api/v1/openprovider/tribonachi/{argument}
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


// Извлекает, валидирует и конвертирует в подходящую
// форму аргумент из пути запроса. Отсекает отрицательные
// значения и значения бльше миллиона (можно убрать проверку
// на верхний предел). В случае не валидного преобразования
// к int - так же возвращает ошибку
func obtainArgument(path string) (int, error) {
	fragments := strings.Split(path,"/")
	lastFragment := fragments[len(fragments)-1]
	argument, err := strconv.Atoi(lastFragment)

	if argument < 0 {
		err = errors.New("negative arguments not supported")
	}

	// можно закомменторовать для стресс-тестов
	if argument > 1000000 {
		err = errors.New("аргументы больше 1000000 запрещены в демонстрационных целях")
	}

	return argument, err
}


// Возвращает значение n-го члена ряда Трибоаччи.
// argument - член ряда Трибоаччи значение которого необходимо получить.
// beginCacheLimit - значение, задающее начальный предел аргумента
//   начиня с котрого вычисленные значения будут кешироваться
func TribonacciThroughCache(argument int, beginCacheLimit int ) string {

	// если аргумент достаточно большой
	if argument > beginCacheLimit {
		//ищем в кеше
		value, found := memoryCache.Get(strconv.Itoa(argument))
		if found {
			return value.(string)
		}
	}

	//при малых значениях аргумента или пустом кеше - вычисляем
	result := TribonacсiIteroBig(argument)

	// если аргумент достаточно большой - запоминаем
	if argument > beginCacheLimit {
		memoryCache.Set(strconv.Itoa(argument), result.String(), cache.DefaultExpiration)
	}

	return  result.String()
}

// Вычисляет значение n-го члена ряда Трибоаччи.
// Используется итеративный алгоритм вычисления.
// Для хранения результатов вычисления используется
// тип даных big.Int не ограничивающий размерность
// (размер переменной типа big.Int ограничен
// только размером оперативной памяти машины).
// Как следствие функция может вычислять значения ряда
// от достаточно  болших значений арумента
// ( миллион и более) за приемлеме время.
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
