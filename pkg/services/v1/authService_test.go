package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateUsertoReturnHttpStatus(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AuthenticateUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	//Positive test
	req.Header.Set("Username", "test1")

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Unauthorized user test
	req, err = http.NewRequest("GET", "/auth", nil)

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(AuthenticateUser)

	req.Header.Set("Username", "zxyzos")

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}
