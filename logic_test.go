package main

import (
	"os"
	"testing"
)

func TestTask(t *testing.T) {
	err := os.WriteFile("todos.json", []byte("[]"), 0644)
	if err != nil {
		t.Fatalf("failed to reset todos.json: %v", err)
	}
	err = AddTask("write test")
	if err != nil {
		t.Errorf("AddTask failed: %v", err)
	}
	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("LoadTasks failed: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Text != "write test" {
		t.Errorf("Task text mismatch: got %q", tasks[0].Text)
	}
}
