package timer

import (
	"testing"
	"time"
	"to-do-list/backend/utils/alert"

	"github.com/stretchr/testify/assert"
)

func TestTimerAlerts(t *testing.T) {
	// Create a new AlertService instance
	alertService := alert.NewAlertService()

	// Create a new TimerService instance
	service := NewService(alertService)

	// Test starting a timer for 6 minutes
	err := service.StartTimer("task-1", 6*time.Minute)
	assert.NoError(t, err)

	// Test stopping the timer
	err = service.StopTimer("task-1")
	assert.NoError(t, err)
}
