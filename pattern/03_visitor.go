package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Посетитель - паттерн поведения объектов.
	Он описывает операцию, выполняемую с каждым объектом из некоторой структуры.
	Паттерн позволяет опеределить новую операцию, не изменяя классы этих объектов.
	Основная идея заключается в том, чтобы вынести операции из классов элементов в отдельные классы посетителей.
	Это позволяет добавлять новые операции, не изменяя сами элементы.

Плюсы:
	Упрощение добавления новых операций.
	Объединение родственных операций и отсечение тех которые не имеют к ним отношения,
Минусы:
	Нарушение инкапсуляции: посетители могут получать доступ к приватным членам элементов, нарушая инкапсуляцию.
	Сложность отладки и понимания кода.
	Повышенная связанность.


Паттер использовать стоит:
1. Когда имеется много объектов разнородных классов с разными интерфейсами, и требуется выполнить ряд операций над каждым из этих объектов.
2. Когда классам необходимо добавить одинаковый набор операций без изменения этих классов.
3. Когда часто добавляются новые операции к классам, при этом общая структура классов стабильна и практически не изменяется.
*/

// Интерфейс элемента.
type Element interface {
	Accept(visitor Visitor)
}

// Элементы.
type Circle struct{}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

type Square struct{}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Интерфейс посетителя. Вычисляет площадь, условно.
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitSquare(square *Square)
}

// Сам посетитель.
type AreaVisitor struct{}

func (av *AreaVisitor) VisitCircle(circle *Circle) {
	/*...*/
}

func (av *AreaVisitor) VisitSquare(square *Square) {
	/*...*/
}

// Есть некая коллекция элементов.
type GeometryStructure struct {
	elements []Element
}

func (gs *GeometryStructure) AddElement(element Element) {
	gs.elements = append(gs.elements, element)
}

// Посетитель проходит по коллекции и выполняет операции к элементу.
func (gs *GeometryStructure) AcceptVisitor(visitor Visitor) {
	for _, element := range gs.elements {
		element.Accept(visitor)
	}
}
