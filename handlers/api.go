package handlers

import (
	"fmt"
	"net/http"

	"github.com/tsukoyachi/goplanner/models"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handle GET Tasks")

	todoList := models.GetTodoList()
	writer.Header().Set("Content-Type", "application/json")
	error := todoList.ToJson(writer)
	if error != nil {
		fmt.Println(todoList)
		http.Error(writer, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func AddTask(responseWriter http.ResponseWriter, request *http.Request) {
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
