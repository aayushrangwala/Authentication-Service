package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	validUser1 = "test1"
	validUser2 = "test2"
)

// AuthenticateUser authenticate the name of the user if present in DB and returns the status code
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["user"]
	if isPresentInDB(userName) {
		json.NewEncoder(w).Encode(strconv.Itoa(http.StatusOK))
	} else {
		json.NewEncoder(w).Encode(strconv.Itoa(http.StatusUnauthorized))
	}
}

func isPresentInDB(u string) bool {
	if u == validUser1 || u == validUser2 {
		return true
	}
	return false
}
