package handlers

import (
	"fmt"
	"net/http"

	"github.com/tsukoyachi/goplanner/models"
)

func TodoHandler(writer http.ResponseWriter, request *http.Request) {
	//TODO: Change with method call instead of print
	switch request.Method {
	case http.MethodGet:
		getTasks(writer)
		return
	case http.MethodPost:
		addTask(writer, request)
		return
	case http.MethodDelete:
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Task deleted")
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(responseWriter http.ResponseWriter) {
	fmt.Println("Handle GET Tasks")

	todoList := models.GetTodoList()
	responseWriter.Header().Set("Content-Type", "application/json")
	error := todoList.ToJson(responseWriter)
	if error != nil {
		fmt.Println(todoList)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func addTask(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Handle POST Task")

	newTask := &models.Task{}
	error := newTask.FromJSON(request.Body)

	if error != nil {
		http.Error(responseWriter, "Failed to unmarshal json", http.StatusBadRequest)
	}

	models.AddTask(newTask)

	responseWriter.WriteHeader(http.StatusCreated)
	error = newTask.ToJson(responseWriter)

	if error != nil {
		fmt.Println(newTask)
		http.Error(responseWriter, "Unable to marshal json but it was added to the todolist", http.StatusInternalServerError)
	}
}
