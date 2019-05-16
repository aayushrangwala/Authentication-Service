package v1

import (
	v1 "Authentication-Service/pkg/services/v1"
	"net/http"
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
