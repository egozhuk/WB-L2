package api

import (
	"time"
)

type Date time.Time

type CreateEventRequest struct {
	UserID      int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        Date   `json:"date"`
}

type UpdateEventRequest struct {
	UserID      int    `json:"user_id"`
	EventID     int    `json:"event_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        Date   `json:"date"`
}

type DeleteEventRequest struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type GetEventsRequest struct {
	UserID int  `json:"user_id"`
	Date   Date `json:"date"`
}

// Преобразуем json строку в тип time
func (d *Date) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	*d = Date(date)
	return nil
}
