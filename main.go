package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/tsukoyachi/goplanner/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todo", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/todo", handlers.AddTask).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/todo/{id}", handlers.DeleteTask).Methods("DELETE")

	// Serve Swagger UI files
	//router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./swagger-ui/"))))

	// Serve Swagger UI
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger"), // URL to your Swagger JSON doc
	))

	// Serve the service.swagger.yaml file
	router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/v1/service.swagger.yml")
	})

	router.HandleFunc("/version", handlers.VersionHandler).Methods("GET")
	router.HandleFunc("/", handlers.GreetingHandler).Methods("GET")

	fmt.Println("The server is up and running on port 8080 !")
	log.Fatal(http.ListenAndServe(":8080", router))
}
