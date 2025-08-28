package main

import (
	"CLI-todo/task"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello CLI-todo! (type help for available commands)")
	manager := task.NewTaskManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "help":
			manager.ListAvailableCommands()
		case "list":
			manager.ListTasks()
		case "add":
			if len(parts) < 4 {
				fmt.Println("Usage: add <title> <description> <priority>")
				continue
			}

			title := parts[1]
			description := strings.Join(parts[2:len(parts)-1], " ")
			priority, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				fmt.Println(parts[3] + " unavailable value for 'priority'")
				continue
			}
			manager.AddTask(title, description, priority)
			fmt.Println("Task added successfully!")
		case "status":
			if len(parts) != 2 {
				fmt.Println("Usage: status <id>")
				continue
			}

			idToChange, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(parts[2] + " unavailable value for 'id'")
				continue
			}

			t, err := manager.ChangeStatus(idToChange)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("status of task '%s' successfully changed to %s\n", t.Title, t.StatusToString())
		case "sort":
			if len(parts) != 2 {
				fmt.Println("Usage: sort <done/undone>")
				continue
			}

			if parts[1] == "done" {
				manager.ListTasksByStatus(true)
				continue
			} else if parts[1] == "undone" {
				manager.ListTasksByStatus(false)
				continue
			} else {
				fmt.Println("Usage: sort <done/undone>")
			}

		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("unknown command")
		}

	}
}
