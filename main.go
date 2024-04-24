package main

import (
	"fmt"
	"log"
	"net/http"
)

func versionHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You're currently using the version 0.0.1")
}

func greetingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, the server is up and running!")
}

func main() {
	http.HandleFunc("/version/", versionHandler)
	http.HandleFunc("/", greetingHandler)
	fmt.Println("The server is up and running on port 8080 !")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
