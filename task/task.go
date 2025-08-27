package task

import "fmt"

var tasks []Task

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    int
	Status      bool
}

func AddTask(title, description string, priority int) Task {
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      false,
	}

	tasks = append(tasks, task)

	return task
}

func ListTasks() {
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Priority: %d, Status: %t\n", task.ID, task.Title, task.Priority, task.Status)
	}
}
