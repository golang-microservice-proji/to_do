package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// to-do list item.
type Task struct {
	ID        string
	Title     string
	Completed bool
	Deadline  time.Time
}

// managing tasks.
type Service struct {
	tasks map[string]*Task
}

// new task service.
func NewService() *Service {
	return &Service{
		tasks: make(map[string]*Task),
	}
}

func (s *Service) CreateTask(title string, deadline time.Time) (*Task, error) {
	id := generateID()
	task := &Task{
		ID:        id,
		Title:     title,
		Completed: false,
		Deadline:  deadline,
	}
	s.tasks[id] = task
	return task, nil
}

func (s *Service) DeleteTask(id string) error {
	_, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}

func (s *Service) MarkTaskComplete(id string) error {
	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	task.Completed = true
	return nil
}

func (s *Service) GetTask(id string) (*Task, error) {
	task, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *Service) ListTasks() []*Task {
	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func generateID() string {
	return uuid.New().String()
}
