package models

import (
	"encoding/json"
	"errors"
	"io"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// This replace the database for the moment
var todoList = Tasks{}

// This will allow us to generate the id each time we'll want a new Task
var sequence int

// This type is a slice of Task
type Tasks []*Task

func (p *Tasks) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)
}

func (p *Task) ToJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(p)
}

func (p *Task) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetTodoList() Tasks {
	return todoList
}

func AddTask(task *Task) {
	task.Id = sequence
	sequence += 1

	todoList = append(todoList, task)
}

func RemoveTask(id int) error {
	for index, value := range todoList {
		if value.Id == id {
			todoList = append(todoList[0:index], todoList[index+1:]...)
			return nil
		}
	}
	return errors.New("id not found")
}

func UpdateTask(task *Task, id int) error {
	for index, value := range todoList {
		if value.Id == id {
			todoList[index].Title = task.Title
			todoList[index].Description = task.Description
			return nil
		}
	}
	return errors.New("id not found")
}
