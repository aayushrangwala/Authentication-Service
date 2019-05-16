package v1

import (
	"net/http"

	v1 "github.com/aayushrangwala/Authentication-Service/pkg/services/v1"
)

// Route struct defines the route mapping
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var route = &Route{
	"AuthenticateUser",
	"GET",
	"/auth/{user}",
	v1.AuthenticateUser,
}
