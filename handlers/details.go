package handlers

import (
	"fmt"
	"net/http"
)

func VersionHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "You're currently using the version 1.0.0")
}

func GreetingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Hi there, the server is up and running!")
}
