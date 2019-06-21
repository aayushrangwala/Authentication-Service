[![Coverage Status](https://coveralls.io/repos/github/aayushrangwala/Authentication-Service/badge.svg?branch=master)](https://coveralls.io/github/aayushrangwala/Authentication-Service?branch=master)


# Authentication-Microservice

This project is a part of a simple application which aims to implement the services required to authenticate and fetch the user details from the database

This Microservice will perform the task of authentication of the user, if it is valid

API Implementations:
  1. API : `/auth` Method : GET, Secured : false, Header {Username: XXXX}
  For every authenticated user request it sends HTTP code 200 to the API gateway and for every un-authenticated request, it sends HTTP status code 401 un-authorized


## How it works?:

The Overall application has the API gateway also as a component which will act as the contact point for the user.
The end user only knows about the two APIS of user-microservice and wi;; request to them. If that request is secured, it will be forwarded to thos service

The Authenticator-Microservice authenticates client HTTP request if the request is valid or not by analyzing the value in the user request header `Username`. If the user with this header is present in the database then it's a valid request otherwise the user request is not a valid request.


## Contents:

Repo contains mainly below files/directories:

1. api/v1:
   1. `routes.go`: Contains the Route struct which is basically the mapping object of the Handler-Function to its url path-pattern
   2. `router.go`: Creates and returns a router object when server is started. Router object contains the list of route objects registerd to it.
2. cmd/server:
   1. `main.go`: It simply creates a router which has mappings of Url path and their handlers and starts a server on port 8080 with it
3. pkg/services/v1:
   1. `authService.go`: This file contains the HandlerFunc for authorizing the Username in the Header. It acts as the main logic of the Microservice
4. `Dockerfile`: It is used to run in the Build and Release Pipeline at Azure infra

## How to Use?

Simply do the `make run` to run the service at the localhost

## Useful Commands:

 1. `make build`: To build the image
 2. `make push`: To push the image
 3. `make test`: To run the tests
 4. `make lint`: To run the lint

## Dependencies

1. [gorilla mux](https://github.com/gorilla/mux)
2. [goveralls](https://github.com/mattn/goveralls)
3. [golangci-lint](https://github.com/golangci/golangci-lint#install)
4. [docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
5. Go installed correctly
