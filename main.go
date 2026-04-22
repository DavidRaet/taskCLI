package main

import (
	"fmt"
	"time"
)

type Status struct {
	todo       string
	inProgress string
	done       bool
}

type Task struct {
	ID        int
	Title     string
	Status    Status
	CreatedAt string
	UpdatedAt string
}

func createTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Status:    Status{todo: "pending"},
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
