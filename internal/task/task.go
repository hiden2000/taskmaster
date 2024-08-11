// internal/task/task.go
package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
}

func NewTask(id int, title, description string, dueDate time.Time) *Task {
	// デフォルトの期限を24時間に設定
	if dueDate.IsZero() {
		dueDate = time.Now().Add(24 * time.Hour)
	}
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) Uncomplete() {
	t.Completed = false
}

func (t *Task) IsOverDue() bool {
	return !t.Completed && time.Now().After(t.DueDate)
}

func (t *Task) String() string {
	status := "Pending"
	if t.Completed {
		status = "Completed"
	} else if t.IsOverDue() {
		status = "Overdue"
	}
	return fmt.Sprintf("Task %d %s [%s] (Due: %s)", t.ID, t.Title, status, t.DueDate.Format(time.RFC822))
}
