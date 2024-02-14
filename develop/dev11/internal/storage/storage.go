package storage

import (
	"log"
	"sync"
	"time"
)

// Event представляет событие в календаре.
type Event struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// Storage представляет хранилище для Event.
type Storage struct {
	storage []Event
	mx      *sync.Mutex
}

// New создает календарь.
func New() *Storage {
	return &Storage{
		storage: make([]Event, 0, 10),
		mx:      &sync.Mutex{},
	}
}

// AddEvent добавляет событие в календарь.
func (s *Storage) AddEvent(event Event) {
	s.mx.Lock()
	s.storage = append(s.storage, event)
	s.mx.Unlock()
}

// DeleteEvent удаляет событие из календаря. Если событие присутствует в календаре, то возвращается true, иначе false.
func (s *Storage) DeleteEvent(event Event) bool {
	isDeleted := false

	s.mx.Lock()

	// Циклом проходит по календарю, и если такое событие есть, то удаляет из календаря.
	for i := 0; i < len(s.storage); i++ {
		if s.storage[i].ID == event.ID {
			s.storage = append(s.storage[:i], s.storage[i+1:]...)
			isDeleted = true
		}
	}

	s.mx.Unlock()

	return isDeleted
}

// UpdateEvent ...
func (s *Storage) UpdateEvent(event Event) {
	log.Println("UpdateEvent")
}

// GetEvents ...
func (s *Storage) GetEvents() {
	log.Println("GetEvents")
}
