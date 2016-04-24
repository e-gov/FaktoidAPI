package rahvafakt

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s", r.Name, r.Method, r.Pattern)
}

type Routes []Route

var routes = Routes{
	Route{
		"faktoid",
		"GET",
		"/faktoid",
		GetFaktoid,
	},
}
