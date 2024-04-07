package alert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertService(t *testing.T) {
	// Create a new AlertService instance
	alertService := NewAlertService()

	// Test registering an alert
	alertChan, err := alertService.RegisterAlert("task-1")
	assert.NoError(t, err)
	assert.NotNil(t, alertChan)

	// Attempt to register the same alert again, expecting an error
	_, err = alertService.RegisterAlert("task-1")
	assert.Error(t, err)

	// Test unregistering an alert
	err = alertService.UnregisterAlert("task-1")
	assert.NoError(t, err)

	// Attempt to unregister a non-existent alert, expecting an error
	err = alertService.UnregisterAlert("task-2")
	assert.Error(t, err)

	// Test sending an alert for a registered task
	err = alertService.SendAlert("task-1")
	assert.Error(t, err) // Expecting an error because the alert was unregistered

	// Register the alert again and test sending an alert
	alertChan, err = alertService.RegisterAlert("task-1")
	assert.NoError(t, err)
	assert.NotNil(t, alertChan)

	go func() {
		// Simulate receiving the alert
		<-alertChan
	}()

	err = alertService.SendAlert("task-1")
	assert.NoError(t, err)

	// Test sending an alert for a non-existent task
	err = alertService.SendAlert("task-2")
	assert.Error(t, err)
}
