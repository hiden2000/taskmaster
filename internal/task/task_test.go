package task

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask(1, "Test Task", "This is a test task", dueDate)

	if task.ID != 1 {
		t.Errorf("Expected task ID 1, got %d", task.ID)
	}
	if task.Title != "Test Task" {
		t.Errorf("Expected task titile 'Test Task', got '%s'", task.Title)
	}
	if task.Completed {
		t.Error("New task should be completed")
	}
}

func TestTaskComplete(t *testing.T) {
	task := NewTask(1, "Test Task", "This is a test task", time.Now())
	task.Complete()

	if !task.Completed {
		t.Error("Task should be completed")
	}
}

func TestTaskIsOverdue(t *testing.T) {
	pastDate := time.Now().Add(-24 * time.Hour)
	task := NewTask(1, "Test Task", "This is test task", pastDate)

	if !task.IsOverDue() {
		t.Error("Task should be overdue")
	}
}
