package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskService(t *testing.T) {
	service := NewService()

	// Test creating a task
	task, err := service.CreateTask("Finish GoLang Proj", time.Now().Add(24*time.Hour))
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Finish GoLang Proj", task.Title)
	assert.False(t, task.Completed)

	// Test getting a task
	retrievedTask, err := service.GetTask(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task, retrievedTask)

	// Test marking a task as complete
	err = service.MarkTaskComplete(task.ID)
	assert.NoError(t, err)
	retrievedTask, err = service.GetTask(task.ID)
	assert.NoError(t, err)
	assert.True(t, retrievedTask.Completed)

	// Test deleting a task
	err = service.DeleteTask(task.ID)
	assert.NoError(t, err)
	_, err = service.GetTask(task.ID)
	assert.Error(t, err)
}
