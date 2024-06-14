package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tsukoyachi/goplanner/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todo", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/todo", handlers.AddTask).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/todo/{id}", handlers.DeleteTask).Methods("DELETE")

	router.HandleFunc("/version", handlers.VersionHandler).Methods("GET")
	router.HandleFunc("/", handlers.GreetingHandler).Methods("GET")

	fmt.Println("The server is up and running on port 8080 !")
	log.Fatal(http.ListenAndServe(":8080", router))
}
