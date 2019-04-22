package router

import "net/http"

// Route is a Routing object
type Route struct {
	Route   string `json:"route"`
	Handler func(http.ResponseWriter, *http.Request)
}

// Router is a Router object
type Router struct {
	Routes   []Route
	Fallback Route
}

// Test is a test
func Test() string {
	return "A"
}
