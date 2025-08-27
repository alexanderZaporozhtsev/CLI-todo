package main

import (
	"CLI-todo/task"
	"fmt"
)

func main() {
	fmt.Println("Hello CLI-todo!")
	task.AddTask("Test title", "Test discription", 5)
	task.ListTasks()
}
