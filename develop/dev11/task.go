package main

import (
	"errors"
	"fmt"
	"log"
	"main/internal/event"
	"main/internal/storage"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

var calendar *storage.Storage

// Обрабатывает запрос на создание события.
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	evt, err := decodeParams(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if validateEvent(evt) {
		fmt.Println("Хорошо")
		calendar.AddEvent(evt)
	} else {
		fmt.Println("Bad")
	}

	log.Println("CreateEventHandler")
}

// Обрабатывает запрос на обновление события.
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")
	description := r.FormValue("description")

	log.Println("UpdateEventHandler", userID, description)
}

// Обрабатывает запрос на удаление события.
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("user_id")

	log.Println("DeleteEventHandler", userID)
}

// Обрабатывает запрос на получение событий за на этот день.
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("events_for_day")
}

// Обрабатывает запрос на получение событий за на эту неделю.
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("events_for_week")
}

// Обрабатывает запрос на получение событий за на этот месяц.
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("events_for_month")
}

// Собирает в структуру полученные из POST-request параметры.
func decodeParams(r *http.Request) (event.Event, error) {
	// Парсит и форматирует под нужный тип.
	eventID, errID := strconv.Atoi(r.FormValue("event_id"))
	if errID != nil {
		return event.Event{}, errors.New(fmt.Sprintf("Неправильный у тебя формат, дружок: %s", errID))
	}

	// Парсит и форматирует под нужный тип.
	date, errDate := time.Parse("2006-01-02", r.FormValue("date"))
	if errDate != nil {
		return event.Event{}, errors.New(fmt.Sprintf("Неправильный у тебя формат, дружок: %s", errDate))
	}

	// Парсит и проверяет на непустоту.
	content := r.FormValue("content")
	if content == "" {
		return event.Event{}, errors.New("Неправильный у тебя формат, дружок: no content!")
	}

	return event.Event{
		ID:      eventID,
		Date:    date,
		Content: content,
	}, nil
}

// Валидация события.
func validateEvent(evnt event.Event) bool {
	return evnt.ID > 0
}

func main() {
	// Подключает обработчики.
	// POST
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)

	// GET
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	calendar = storage.New()

	// Запускает сервер.
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
