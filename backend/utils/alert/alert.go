package alert

import (
	"fmt"
	"sync"
	"time"
)

// AlertService provides functionality for generating task alerts.
type AlertService struct {
	alertChannels map[string]chan bool
	mu            sync.Mutex
}

// NewAlertService creates a new AlertService instance.
func NewAlertService() *AlertService {
	return &AlertService{
		alertChannels: make(map[string]chan bool),
	}
}

// RegisterAlert registers an alert channel for a task.
func (a *AlertService) RegisterAlert(taskID string) (chan bool, error) {
	if _, ok := a.alertChannels[taskID]; ok {
		return nil, fmt.Errorf("alert already registered for task %s", taskID)
	}

	alertChan := make(chan bool)
	a.alertChannels[taskID] = alertChan
	return alertChan, nil
}

// UnregisterAlert unregisters an alert channel for a task.
func (a *AlertService) UnregisterAlert(taskID string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	alertChan, ok := a.alertChannels[taskID]
	if !ok {
		return fmt.Errorf("no alert registered for task %s", taskID)
	}

	close(alertChan)
	delete(a.alertChannels, taskID)
	return nil
}

// SendAlert sends an alert for the given task.
func (a *AlertService) SendAlert(taskID string) error {
	alertChan, ok := a.alertChannels[taskID]
	if !ok {
		return fmt.Errorf("no alert registered for task %s", taskID)
	}

	select {
	case alertChan <- true:
		return nil
	case <-time.After(1 * time.Second):
		return fmt.Errorf("failed to send alert for task %s", taskID)
	}
}
