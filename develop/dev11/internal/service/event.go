package service

import (
	"WB-L2/develop/dev11/internal/domain"
	"WB-L2/develop/dev11/internal/repository"
	"context"
	"fmt"
	"time"
)

type event struct {
	repo repository.Repository
}

func NewEvent(repo repository.Repository) Event {
	return &event{repo: repo}
}

// Получаем событие по дню
func (e *event) GetEventsForDay(ctx context.Context, userID int, day time.Time) ([]domain.Event, error) {
	events, err := e.repo.GetEvents(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []domain.Event
	for _, event := range events {
		if event.Date.Year() == day.Year() && event.Date.YearDay() == day.YearDay() {
			result = append(result, event)
		}
	}
	return result, nil
}

// Получаем событие по неделе
func (e *event) GetEventsForWeek(ctx context.Context, userID int, week time.Time) ([]domain.Event, error) {
	events, err := e.repo.GetEvents(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []domain.Event
	year, weekNumber := week.ISOWeek()
	for _, event := range events {
		eventYear, eventWeek := event.Date.ISOWeek()
		if eventYear == year && eventWeek == weekNumber {
			result = append(result, event)
		}
	}
	return result, nil
}

// Получаем событие по месяцу
func (e *event) GetEventsForMonth(ctx context.Context, userID int, month time.Time) ([]domain.Event, error) {
	events, err := e.repo.GetEvents(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []domain.Event
	for _, event := range events {
		if event.Date.Year() == month.Year() && event.Date.Month() == month.Month() {
			result = append(result, event)
		}
	}
	return result, nil
}

// Создать новый ивент
func (e *event) CreateEvent(ctx context.Context, userID int, event domain.Event) error {
	return e.repo.CreateEvent(ctx, userID, event)
}

// Обновить ивент
func (e *event) UpdateEvent(ctx context.Context, userID int, event domain.Event) error {
	events, err := e.repo.GetEvents(ctx, userID)
	if err != nil {
		return err
	}
	for i, e := range events {
		if e.ID == event.ID {
			events[i] = event
			return nil
		}
	}
	return fmt.Errorf("event %d not found for user %d", event.ID, userID)
}

// Удалить ивент
func (e *event) DeleteEvent(ctx context.Context, userID int, eventID int) error {
	return e.repo.DeleteEvent(ctx, userID, eventID)
}
