package handlers

import (
	"fmt"
	"net/http"
)

func VersionHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You're currently using the version 0.0.1")
}

func GreetingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, the server is up and running!")
}