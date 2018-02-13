package openprovider

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


// Name - Имя запроса.
// Method - Http метод запроса (GET, POST, PUT, DELETE, ...)
// Pattern - Entrypoint запроса (путь + аргументы)
// HandlerFunc - Имя функции-бработчика запроса.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Обработчик Entrypoints.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Entrypoint.
// Обработка запроса к /api/v1/openprovider
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Openprovider!")
}

// Достуные Enrypoints.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/openprovider/",
		Index,
	},

	Route{
		"GetTribonacсiValue",
		"GET",
		"/api/v1/openprovider/tribonachi/{argument}",
		GetTribonacсiValue,
	},
}
