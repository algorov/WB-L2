package event

import "time"

// Event представляет событие в календаре.
type Event struct {
	ID      int       `json:"event_id"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}
