package timer

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	"to-do-list/backend/utils/alert"
)

type Service struct {
	timers       map[string]*Timer
	alertService *alert.AlertService
	mu           sync.Mutex
}

// timer for a task.
type Timer struct {
	TaskID    string
	StartTime time.Time
	Duration  time.Duration
	AlertChan chan bool
}

func NewService(alertService *alert.AlertService) *Service {
	return &Service{
		timers:       make(map[string]*Timer),
		alertService: alertService,
	}
}

func (s *Service) StartTimer(taskID string, duration time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.timers[taskID]
	if ok {
		return errors.New("timer already started for the task")
	}

	alertChan, err := s.alertService.RegisterAlert(taskID)
	if err != nil {
		return err
	}

	timer := &Timer{
		TaskID:    taskID,
		StartTime: time.Now(),
		Duration:  duration,
		AlertChan: alertChan,
	}
	s.timers[taskID] = timer

	go s.runTimer(timer)
	return nil
}

func (s *Service) StopTimer(taskID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.timers[taskID]
	if !ok {
		return errors.New("no timer found for the task")
	}

	err := s.alertService.UnregisterAlert(taskID)
	if err != nil {
		return err
	}

	delete(s.timers, taskID)
	return nil
}

func (s *Service) runTimer(timer *Timer) {
	// Calculate the time at which the first alert should be sent (5 minutes before the timer's duration expires)
	alertTime := timer.StartTime.Add(timer.Duration - 5*time.Minute)

	for {
		// Check if the current time is after the timer's start time plus its duration
		if time.Now().After(timer.StartTime.Add(timer.Duration)) {
			// Timer has expired, send an alert saying "time is up" and remove the timer
			err := s.alertService.SendAlert(timer.TaskID)
			if err != nil {
				log.Printf("Failed to send alert for task %s: %v", timer.TaskID, err)
			}
			s.StopTimer(timer.TaskID)
			return
		}

		// Check if the current time is after the alertTime
		if time.Now().After(alertTime) {
			// 5 minutes left, send an alert saying "5 minutes remaining"
			err := s.alertService.SendAlert(timer.TaskID)
			if err != nil {
				log.Printf("Failed to send alert for task %s: %v", timer.TaskID, err)
			}
			// Recalculate alertTime to be 5 minutes before the timer's expiration time
			alertTime = timer.StartTime.Add(timer.Duration - 5*time.Minute)
		}

		// Sleep for a short duration to avoid busy-waiting
		time.Sleep(1 * time.Second)
	}
}

func (s *Service) GetAlert(taskID string) (chan bool, error) {
	timer, ok := s.timers[taskID]
	if !ok {
		return nil, fmt.Errorf("no timer found for task %s", taskID)
	}
	return timer.AlertChan, nil
}
