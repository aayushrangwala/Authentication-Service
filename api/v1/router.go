package v1

import (
	"Authentication-Service/util"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the list of router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	var handler http.Handler

	handler = route.HandlerFunc
	handler = util.Logger(handler, route.Name)
	router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)

	return router
}
