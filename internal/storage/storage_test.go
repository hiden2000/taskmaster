// internal/storage/storage_test.go

package storage

import (
	"testing"
	"time"

	"github.com/hiden2000/taskmaster/internal/task"
)

func TestAddAndGetTask(t *testing.T) {
	s := NewStorage() // IDの割り振りは 1 からスタート
	newTask := task.NewTask(0, "Test Task", "This is a test task", time.Now())

	err := s.AddTask(newTask)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	retrievedTask, err := s.GetTask(1)
	if err != nil {
		t.Fatalf("Failed to get task: %v", err)
	}

	if retrievedTask.Title != newTask.Title {
		t.Errorf("Expected task title %s, got %s", newTask.Title, retrievedTask.Title)
	}
}

func TestUpdateTask(t *testing.T) {
	s := NewStorage()
	newTask := task.NewTask(0, "Test Task", "This is a test task", time.Now())

	s.AddTask(newTask)

	updatedTask := task.NewTask(1, "Updated Task", "This is an updated task", time.Now())
	err := s.UpdateTask(updatedTask)
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	retrievedTask, _ := s.GetTask(1)
	if retrievedTask.Title != "Updated Task" {
		t.Errorf("Expected updated task title 'Updated Task', got '%s'", retrievedTask.Title)
	}
}

func TestDeleteTask(t *testing.T) {
	s := NewStorage() // 同じくIDの割り振りは 1 からスタートすることに注意
	newTask := task.NewTask(0, "Test Task", "This is a test task", time.Now())

	s.AddTask(newTask)

	err := s.DeleteTask(1)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	_, err = s.GetTask(1)
	if err == nil {
		t.Error("Expected error when getting deleted task, got nil")
	}
}
