package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("Usage: go run main.go <command> [arguments]")
		return
	}

	command := arg[1]
	switch command {
	case "add":
		if len(arg) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		text := arg[2]
		for i := 3; i < len(arg); i++ {
			text += arg[i]
		}
		err := AddTask(text)
		if err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			fmt.Println("Task added successfully ✅")
		}
	case "list":
		err := ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}

	case "done":
		if len(arg) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		index, err := strconv.Atoi(arg[2])
		if err != nil {
			fmt.Println("Invalid task number.")
			return
		}
		err = DoneTask(index)
		if err != nil {
			fmt.Println("Error marking task done:", err)
		}
	case "delete":
		if len(arg) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		index, err := strconv.Atoi(arg[2])
		if err != nil {
			fmt.Println("Invalid task number.")
			return
		}
		err = DeleteTask(index)
		if err != nil {
			fmt.Println("Error deleting task:", err)
		}
	case "filter":
		filter := ""
		if len(arg) > 2 {
			filter = arg[2]
			err := ListFilteredTasks(filter)
			if err != nil {
				fmt.Println("Error listing filtered tasks:", err)
			}
			return
		}

		err := ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: add")
	}

}
