package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Task struct {
	Text string
	Done bool
}

// Read from todos.json
func LoadTasks() ([]Task, error) {
	file, err := os.Open("todos.json")
	// explicit error checking
	if err != nil {
		if os.IsNotExist(err) {
			// Prompt the user
			fmt.Print("No todo file found. Do you want to create one? (y/n): ")
			var response string
			fmt.Scanln(&response)
			if response == "y" || response == "Y" {
				// Create empty file with []
				err := os.WriteFile("todos.json", []byte("[]"), 0644)
				if err != nil {
					return nil, err
				}
				fmt.Println("Created new todos.json ")
				return []Task{}, nil
			} else {
				return nil, fmt.Errorf("todo file not found and user declined to create one")
			}
		}

		return nil, err
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		// File exists but is empty → treat as empty JSON array
		data = []byte("[]")
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}

// Write tasks to todos.json
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		return err
	}

	return nil

}

func AddTask(task string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		Text: task,
		Done: false,
	}
	tasks = append(tasks, newTask)

	return SaveTasks(tasks)

}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return nil
	}
	for i, task := range tasks {
		var status string
		if task.Done {
			status = "DONE   "
		} else {
			status = "PENDING"
		}
		fmt.Printf("%d. %-8s %s\n", i+1, status, task.Text)

	}
	return nil
}

func DoneTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("invalid task number: %d", index)
	}

	tasks[index-1].Done = true

	err = SaveTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf(" Task %d marked as done!\n", index)
	return nil
}

func DeleteTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	if index < 1 || index > len(tasks) {
		return fmt.Errorf("invalid task number: %d", index)
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	err = SaveTasks(tasks)
	if err != nil {
		return err
	}
	fmt.Printf("🗑️ Task %d deleted!\n", index)
	return nil
}

func ListFilteredTasks(filter string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return nil
	}

	for i, task := range tasks {
		if filter == "done" && !task.Done {
			continue
		}
		if filter == "pending" && task.Done {
			continue
		}

		status := "PENDING"
		if task.Done {
			status = "DONE   "
		}
		fmt.Printf("%d. %-8s %s\n", i+1, status, task.Text)
	}

	return nil
}
