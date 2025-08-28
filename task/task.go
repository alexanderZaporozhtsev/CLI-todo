package task

import (
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    int
	Status      bool
}

type Command struct {
	Name        string
	Description string
}

func (t Task) StatusToString() string {
	if t.Status {
		return "[✔]"
	}
	return "[ ]"
}

type TaskManager struct {
	tasks    []Task
	commands []Command
}

func (m *TaskManager) ListAvailableCommands() {
	fmt.Println("Available commands:")
	fmt.Println("-------------------")

	for _, cmd := range m.commands {
		fmt.Printf("• %-10s : %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println("-------------------")
}

func (m *TaskManager) initCommands() {
	m.commands = []Command{
		{Name: "add", Description: "Add a new task"},
		{Name: "list", Description: "List all tasks"},
		{Name: "sort", Description: "List tasks by status"},
		{Name: "status", Description: "Change task status"},
		{Name: "delete", Description: "Delete a task"},
		{Name: "help", Description: "Show all commands"},
		{Name: "exit", Description: "Quite programm"},
	}
}

func (m *TaskManager) ListTasksByStatus(status bool) {
	if len(m.tasks) == 0 {
		fmt.Println("no tasks available")
		return
	}

	for i := range m.tasks {
		if m.tasks[i].Status == status {
			status := m.tasks[i].StatusToString()

			fmt.Printf("ID: %d\n Status: %s\n Priority: %d\n Title: %s\n Description: %s\n",
				m.tasks[i].ID, status, m.tasks[i].Priority, m.tasks[i].Title, m.tasks[i].Description)
		}
	}
}

func (m *TaskManager) DeleteTask(id int) error {
	for i := range m.tasks {
		if id == m.tasks[i].ID {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID: %d not found", id)
}

func (m *TaskManager) AddTask(title, description string, priority int) {
	id := len(m.tasks) + 1
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      false,
	}

	m.tasks = append(m.tasks, task)
}

func (m *TaskManager) ChangeStatus(id int) (Task, error) {
	for i := range m.tasks {
		if id == m.tasks[i].ID {
			m.tasks[i].Status = !m.tasks[i].Status
			return m.tasks[i], nil
		}
	}

	return Task{}, fmt.Errorf("task with ID: %d not found", id)
}

func (m *TaskManager) ListTasks() {
	if len(m.tasks) == 0 {
		fmt.Println("list is empty")
		return
	}

	for _, t := range m.tasks {
		status := t.StatusToString()

		fmt.Printf("ID: %d\n Status: %s\n Priority: %d\n Title: %s\n Description: %s\n",
			t.ID, status, t.Priority, t.Title, t.Description)
	}
}

func NewTaskManager() *TaskManager {
	m := &TaskManager{
		tasks: []Task{},
	}

	m.initCommands()

	return m
}
