package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsukoyachi/goplanner/handlers"
)

func main() {
	http.HandleFunc("/todo", handlers.TodoHandler)
	http.HandleFunc("/version/", handlers.VersionHandler)
	http.HandleFunc("/", handlers.GreetingHandler)
	fmt.Println("The server is up and running on port 8080 !")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
