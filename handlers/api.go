package handlers

import (
	"fmt"
	"net/http"
)

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO: Change with method call instead of print
	switch request.Method {
	case http.MethodGet:
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "This is your TODO list")
	case http.MethodPost:
		writer.WriteHeader(http.StatusCreated)
		fmt.Fprintf(writer, "Task created")
	case http.MethodDelete:
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Task deleted")
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
