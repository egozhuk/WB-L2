package inmemory

import (
	"WB-L2/develop/dev11/internal/domain"
	"context"
	"fmt"
	"sync"
)

type event struct {
	mutex sync.RWMutex
	data  map[int][]domain.Event
}

func NewEvent() Event {
	return &event{data: make(map[int][]domain.Event)}
}

// Добавить ивент
func (e *event) CreateEvent(ctx context.Context, userID int, event domain.Event) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	event.ID = len(e.data[userID]) + 1
	e.data[userID] = append(e.data[userID], event)
	return nil
}

// Солучить события по юзеру
func (e *event) GetEvents(ctx context.Context, userID int) ([]domain.Event, error) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	events, exists := e.data[userID]
	if !exists || len(events) == 0 {
		return nil, fmt.Errorf("no events found for user %d", userID)
	}
	return events, nil
}

// Удалить событие
func (e *event) DeleteEvent(ctx context.Context, userID int, eventID int) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	events, exists := e.data[userID]
	if !exists {
		return fmt.Errorf("no events found for user %d", userID)
	}
	for i, event := range events {
		if event.ID == eventID {
			e.data[userID] = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("event %d not found for user %d", eventID, userID)
}
