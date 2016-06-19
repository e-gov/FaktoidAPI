package faktoid

import (
	"fmt"
	"net/http"
)

// Route relates parameters of a HTTP request to a handler function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s", r.Name, r.Method, r.Pattern)
}

// Routes contains the currently active routes
// In this case, the standard FaktoidAPI URLs and standard handlers
type Routes []Route

var routes = Routes{
	Route{
		"faktoid",
		"GET",
		"/faktoid",
		GetFaktoid,
	},
	Route{
		"filteredfaktoid",
		"GET",
		"/faktoid/{filter}",
		GetFilteredFaktoid,
	},
	Route{
		"data",
		"GET",
		"/andmed",
		GetData,
	},
	Route{
		"meta",
		"GET",
		"/meta",
		GetMeta,
	},
}
