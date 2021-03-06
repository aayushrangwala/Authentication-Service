package v1

import (
	"net/http"
)

const (
	validUser1 = "test1"
	validUser2 = "test2"
)

// AuthenticateUser authenticate the name of the user if present in DB and returns the status code
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	if isPresentInDB(username) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func isPresentInDB(u string) bool {
	if u == validUser1 || u == validUser2 {
		return true
	}
	return false
}
