package main

import (
	"log"
	"net/http"

	v1 "github.com/aayushrangwala/Authentication-Service/api/v1"
)

func main() {

	router := v1.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
