package main

import (
	v1 "Authentication-Service/api/v1"
	"log"
	"net/http"
)

func main() {

	router := v1.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
