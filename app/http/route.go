package route

import (
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/gorilla/mux"
	"github.com/iammallik/sample-heroku-go/app/http/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},

	Route{
		"CheckLiveness",
		"GET",
		"/liveness",
		handlers.Liveness,
	},

	Route{
		"CheckReadiness",
		"GET",
		"/readiness",
		handlers.Readiness,
	},
}

//Attaches debug profiler to routes.
func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
}

//Adds endpoints and its handlers to router.
func addRoutes(router *mux.Router) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

//Creates a new router with needed route and profiler setting
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.KeepContext = true

	if os.Getenv("DEBUG_PROFILE") == "enable" {
		attachProfiler(router)
	}

	addRoutes(router)

	return router
}
