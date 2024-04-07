package models

import (
	"time"
)

// Task represents a to-do list item.
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Deadline  time.Time `json:"deadline"`
}
