package storage

import (
	"main/internal/event"
	"sync"
	"time"
)

// Storage представляет хранилище для Event.
type Storage struct {
	Storage []event.Event
	mx      *sync.Mutex
}

// New создает календарь.
func New() *Storage {
	return &Storage{
		Storage: make([]event.Event, 0, 10),
		mx:      &sync.Mutex{},
	}
}

// AddEvent добавляет событие в календарь.
func (s *Storage) AddEvent(event event.Event) {
	s.mx.Lock()
	s.Storage = append(s.Storage, event)
	s.mx.Unlock()
}

// DeleteEvent удаляет событие из календаря. Если событие присутствует в календаре, то возвращается true, иначе false.
func (s *Storage) DeleteEvent(eventID int) bool {
	isDeleted := false

	s.mx.Lock()

	// Циклом проходит по календарю, и если такое событие есть, то удаляет из календаря.
	for i := 0; i < len(s.Storage); i++ {
		if s.Storage[i].ID == eventID {
			s.Storage = append(s.Storage[:i], s.Storage[i+1:]...)
			isDeleted = true
		}
	}

	s.mx.Unlock()

	return isDeleted
}

// UpdateEvent обновляет описание события.
func (s *Storage) UpdateEvent(eventID int, content string) bool {
	flag := false

	s.mx.Lock()

	// Циклом проходит по календарю, и если такое событие есть, то удаляет из календаря.
	for i := 0; i < len(s.Storage); i++ {
		if s.Storage[i].ID == eventID {
			s.Storage[i].Content = content
			flag = true
		}
	}

	s.mx.Unlock()

	return flag
}

// ByDay возвращает слайс событий по дню.
func (s *Storage) ByDay(date time.Time) []event.Event {
	s.mx.Lock()

	var events []event.Event

	for _, e := range s.Storage {
		// Если год, месяц и день идентичны, то добавляет в результирующий слайс.
		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() && e.Date.Day() == date.Day() {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}

// ByWeek возвращает слайс событий по неделе.
func (s *Storage) ByWeek(date time.Time) []event.Event {
	s.mx.Lock()

	var events []event.Event

	for _, e := range s.Storage {
		// Рассчитывает абсолютную разность между двумя датами.
		difference := date.Sub(e.Date)
		if difference < 0 {
			difference = -difference
		}

		// Если разность меньше или равно неделе, то добавляет в слайс.
		if difference <= time.Duration(7*24)*time.Hour {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}

// ByMonth возвращает слайс событий по месяцу.
func (s *Storage) ByMonth(date time.Time) []event.Event { //список событий с совпадением дат вплоть до месяца
	s.mx.Lock()

	var events []event.Event

	for _, e := range s.Storage {
		// Если месяцы равно, то они идентично подходят.
		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() {
			events = append(events, e)
		}
	}

	s.mx.Unlock()

	return events
}
