// internal/storage/storage.go

package storage

import (
	"fmt"
	"sync"

	"github.com/hiden2000/taskmaster/internal/task"
)

type Storage struct {
	tasks  map[int]*task.Task
	mutex  sync.RWMutex
	nextID int
}

func NewStorage() *Storage {
	return &Storage{
		tasks:  make(map[int]*task.Task),
		nextID: 1,
	}
}

func (s *Storage) AddTask(t *task.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	t.ID = s.nextID
	s.tasks[t.ID] = t
	s.nextID++
	return nil
}

func (s *Storage) GetTask(id int) (*task.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	t, exists := s.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task with id %d not found", id)
	}
	return t, nil
}

func (s *Storage) UpdateTask(t *task.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.tasks[t.ID]; !exists {
		return fmt.Errorf("task with id %d not found", t.ID)
	}
	s.tasks[t.ID] = t
	return nil
}

func (s *Storage) DeleteTask(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("task with id %d not found", id)
	}
	delete(s.tasks, id)
	return nil
}

func (s *Storage) ListTasks() []*task.Task {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks := make([]*task.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}
