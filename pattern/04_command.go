package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Используется, к примеру, в качестве кнопок для отдельных действий, задание макросов, многоуровневой отмены операций, индикаторы выполнения, транзакции и др.

	Плюсы:
		1. Отделение клиента от получателя:
			* Клиент может работать с командой, не зная, как именно выполняется операция.
			* Это позволяет изменять и добавлять операции, не затрагивая клиентский код.
		2. Поддержка отмены и повторения операций:
			* Так как команда представлена в виде объекта, вы можете легко реализовать механизмы отмены и повторения операций.
		3. Хранение истории операций:
			* Команды могут быть сохранены в истории, что позволяет воспроизводить последовательность операций или создавать снимки состояния системы.
	Минусы:
		1. Увеличение числа классов.
		2. Повышенное использование памяти.
		3. Сложность реализации отмены и повторения.
*/

// Command интерфейс определяет метод Execute.
type Command interface {
	Execute()
}

// LightReceiver - получатель команды (объект, над которым выполняется действие).
type LightReceiver struct{}

func (l *LightReceiver) TurnOn() {
	fmt.Println("Свет включен")
}

func (l *LightReceiver) TurnOff() {
	fmt.Println("Свет выключен")
}

// On - конкретная команда для включения света.
type On struct {
	Light *LightReceiver
}

func (c *On) Execute() {
	c.Light.TurnOn()
}

// Off - конкретная команда для выключения света.
type Off struct {
	Light *LightReceiver
}

func (c *Off) Execute() {
	c.Light.TurnOff()
}

// Invoker - структура, которая вызывает команды.
type Invoker struct {
	Command Command
}

func (i *Invoker) PressButton() {
	i.Command.Execute()
}

func main() {
	// Создаем объект получателя (света).
	light := &LightReceiver{}

	// Создаем конкретные команды с объектом получателя.
	commandOn := &On{Light: light}
	commandOff := &Off{Light: light}

	// Создаем инициатор с командами.
	switchOnButton := &Invoker{Command: commandOn}
	switchOffButton := &Invoker{Command: commandOff}

	// Вызываем команды через инициатор.
	switchOnButton.PressButton()
	switchOffButton.PressButton()
}
