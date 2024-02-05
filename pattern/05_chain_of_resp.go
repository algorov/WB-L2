package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Chain of Responsibility — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый обработчик решает, может ли он обработать запрос, и передает его следующему обработчику в цепочке, если не может.

Такой подход делает систему более гибкой и расширяемой, так как можно легко добавлять или изменять обработчики без изменения кода клиента.

Плюсы:
	* Гибкость и расширяемость;
	* Отделение клиента от обработчиков;
	* Единообразие запросов.

Минусы:
	* Нет гарантии обработки;
	* При бОльшем количестве обработчиков, возможна неэффективность.
	* Сложность откладки.
*/

package main

import "fmt"

// Объявление интерфейса обработчика.
type Handler interface {
	handleRequest(request int)
	setNextHandler(handler Handler)
}

// Реализация интерфейса + встраивание.
type ConcreteHandler struct {
	nextHandler Handler
}

// Реализация контракта.
func (ch *ConcreteHandler) setNextHandler(handler Handler) {
	ch.nextHandler = handler
}

// Реализация контракта.
func (ch *ConcreteHandler) handleRequest(request int) {
	if request <= 10 {
		fmt.Printf("Обработчик обрабатывает запрос %d\n", request)
	} else if ch.nextHandler != nil {
		fmt.Println("Обработчик передает запрос следующему")
		ch.nextHandler.handleRequest(request)
	} else {
		fmt.Println("Нет подходящего обработчика, запрос не обработан.")
	}
}

func main() {
	// Создание обработчиков
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	// Формирование цепочки обработчиков
	handler1.setNextHandler(handler2)
	handler2.setNextHandler(handler3)

	// Передача запросов по цепочке
	handler1.handleRequest(5)
	handler1.handleRequest(12)
}
