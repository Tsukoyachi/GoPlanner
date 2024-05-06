package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tsukoyachi/goplanner/handlers"
)

func main() {
	http.HandleFunc("GET /todo", handlers.GetTasks)
	http.HandleFunc("POST /todo", handlers.AddTask)
	http.HandleFunc("GET /version", handlers.VersionHandler)
	http.HandleFunc("GET /", handlers.GreetingHandler)
	fmt.Println("The server is up and running on port 8080 !")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
