package main

import (
	"fmt"
	"time"
	"os"
	"encoding/json"
)

type Status struct {
    Todo       string `json:"todo"`
    InProgress string `json:"inProgress"`
    Done       bool   `json:"done"`
}

type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Status    Status `json:"status"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

func createTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Status:    Status{Todo: "pending"},
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
}

func main() {
	newTask := createTask(123, "New Task")

	fmt.Printf("Hello new task: %s\n", newTask.Title)
	fmt.Printf("New task ID: %d\n", newTask.ID)
	fmt.Printf("Task status: %s\n", newTask.Status.todo)
	fmt.Printf("Task created at: %s\n", newTask.CreatedAt)
}
