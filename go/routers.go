package openprovider

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Name - Name of request.
// Method - Http method of request (GET, POST, PUT, DELETE, ...)
// Pattern - Entrypoint of request (path + args)
// HandlerFunc - Function-handler name of request.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Handler of Entrypoints.
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

// Entrypoint /api/v1/openprovider
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Openprovider!")
}

// Available Enrypoints.
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
