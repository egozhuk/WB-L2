package service

import (
	"WB-L2/develop/dev11/internal/domain"
	"WB-L2/develop/dev11/internal/repository"
	_ "WB-L2/develop/dev11/internal/repository"
	"context"
	"time"
)

// Сервис ивентов
type Event interface {
	GetEventsForDay(ctx context.Context, userIdD int, day time.Time) ([]domain.Event, error)
	GetEventsForWeek(ctx context.Context, userID int, week time.Time) ([]domain.Event, error)
	GetEventsForMonth(ctx context.Context, userID int, month time.Time) ([]domain.Event, error)
	CreateEvent(ctx context.Context, userID int, event domain.Event) error
	UpdateEvent(ctx context.Context, userID int, event domain.Event) error
	DeleteEvent(ctx context.Context, userID int, eventID int) error
}

// Набор всех сервисов
type Service struct {
	Event
}

func NewService(repos repository.Repository) Service {
	return Service{Event: NewEvent(repos)}
}
