package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func DeleteTask(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Handle DELETE Task")

	vars := mux.Vars(request)
	idString, ok := vars["id"]
	if !ok {
		http.Error(responseWriter, "id is missing in parameters", http.StatusBadRequest)
	}

	id, error := strconv.Atoi(idString)
	if error != nil {
		http.Error(responseWriter, "id is not a number", http.StatusBadRequest)
	}

	error = models.RemoveTask(id)

	if error != nil {
		http.Error(responseWriter, "The task was not found", http.StatusNotFound)
	} else {
		io.WriteString(responseWriter, "The task was deleted successfully !")
	}
}

func UpdateTask(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Handle PUT Task")

	newTask := &models.Task{}
	error := newTask.FromJSON(request.Body)

	if error != nil {
		http.Error(responseWriter, "Failed to unmarshal json", http.StatusBadRequest)
	}

	vars := mux.Vars(request)
	idString, ok := vars["id"]
	if !ok {
		http.Error(responseWriter, "id is missing in parameters", http.StatusBadRequest)
	}

	id, error := strconv.Atoi(idString)
	if error != nil {
		http.Error(responseWriter, "id is not a number", http.StatusBadRequest)
	}

	error = models.UpdateTask(newTask, id)
	if error != nil {
		http.Error(responseWriter, "The task was not found", http.StatusNotFound)
	}

	newTask.Id = id

	responseWriter.WriteHeader(http.StatusOK)
	error = newTask.ToJson(responseWriter)

	if error != nil {
		fmt.Println(newTask)
		http.Error(responseWriter, "Unable to marshal json but it was updated in the todolist", http.StatusInternalServerError)
	}
}
