package main

import (
	"CLI-todo/task"
	"bufio"
	"fmt"
	"os"
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

		switch parts[0] {
		case "help":
			manager.ListAvailableCommands()
		case "list":
			manager.ListTasks()
		default:
			fmt.Println("unknown command")
		}

	}
}
