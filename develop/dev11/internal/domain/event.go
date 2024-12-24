package domain

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Date        time.Time
}
