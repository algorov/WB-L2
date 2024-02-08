package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Поведенческий паттерн проектирования. Основан на алгоритме для определения соответсвущего класса.
Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.
Паттерн отделяет процедуру выбора конкретного воедения от реализации, что позволяет.

Нужда:
* Программа должна обеспечивать различные варианты алгоритма или поведения.
* Нужно изменять поведение каждого экземпляра класса.
* Необходимо изменять поведение объектов на стадии выполнения.
* Введение интерфейса позволяет классам-клиентам ничего не знать о классах, реализующих этот интерфейс и инкапсулирующих в себе конкретные алгоритмы.

Применяемость:
Допустим, есть игрок и несколько его поведений, на основе какого-то действия одно поведение заменяется на другое.
То есть игроку не важно, какое это конкретное поведение, поскольку реализация его индивидуальна (по-свему реализует интерфейс).

Плюсы:
Динамическое определение, какой алгоритм будет запущен.
Инкапсуляция. Отделен код конкретной стратегии от остального кода.

Минусы:
Нагроможденность кода из-за дополнительных структур. Также нужно четко понимать, что (статегию) и где применять.

*/

// Интерфейс стратегии.
type Move interface {
	Exec()
}

// Конкретная стратегия.
type MoveX struct {
	// *** //
}

// Реализация конкретной стратегии.
func (mx MoveX) Exec() {
	// *** //
}

// Конкретная стратегия.
type MoveY struct {
	// *** //
}

// Реализация конкретной стратегии.
func (my MoveY) Exec() {
	// *** //
}

type Object struct {
	// *** //
}

func (o *Object) Go(strategy Move) {
	strategy.Exec()
}
