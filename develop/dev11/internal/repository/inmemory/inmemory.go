package inmemory

import (
	"WB-L2/develop/dev11/internal/domain"
	"context"
)

// Хранилище интерфейса
type Event interface {
	CreateEvent(ctx context.Context, userID int, event domain.Event) error
	GetEvents(ctx context.Context, userID int) ([]domain.Event, error)
	DeleteEvent(ctx context.Context, userID int, eventID int) error
}

// Хранилище локальное
type InMemory struct {
	Event
}

func NewInMemory() InMemory {
	return InMemory{Event: NewEvent()}
}
