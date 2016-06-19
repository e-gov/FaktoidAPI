package faktoid

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// NewRouter returns a properly initialized router with request handlers and logging attached
func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		
		handler = LoggingHandler(handler, log)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handlers.LoggingHandler(os.Stdout, handler))
		log.Debug("Added " + route.String())
		
	}
	return router
}
