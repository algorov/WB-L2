package dev02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unwrap(str string) (string, error) {
	// Создает билдер, чтобы эффективно конкатенировать строки.
	builder := strings.Builder{}

	// Перевод полученной строки в слайс рун.
	runes := []rune(str)

	// Рассматриваемая руна, которая будет множиться в зависимости от коэффециента.
	var procRune *rune

	// Итерация по срезу.
	for i := 0; i < len(runes); i++ {
		// Для удобства отдельная переменная, которая хранит текущую руну.
		curRune := runes[i]

		// Если нашлелся символ экранирования.
		if curRune == '\\' {
			// Проверка, не последний ли этот элемент в слайсе.
			if i+1 > len(runes)-1 {
				return "", errors.New("Invalid format")
			}

			// Добавление экранируемого элемента.
			builder.WriteRune(runes[i+1])
			procRune = new(rune)
			*procRune = runes[i+1]

			// Сдвиг указателя.
			i++

			// Переход на следующую итерацию.
			continue
		}

		// Если текущая руна не цифра.
		if !unicode.IsDigit(curRune) {
			// Запись в билдер.
			builder.WriteRune(curRune)

			// Установка руны как обрабатываемого.
			procRune = new(rune)
			*procRune = runes[i]
		} else {
			// Если текущая руна - цирфа,
			// Проверка, что эта руна служит коэффициентом для обрабатываемой руны.
			if procRune != nil {
				// Перевод руны в число.
				count, err := strconv.Atoi(string(curRune))
				if err != nil {
					return "", err
				}

				// Множение руны в N раз.
				for i := 1; i < count; i++ {
					builder.WriteRune(*procRune)
				}

				// Обработанную руну сбрасывает.
				procRune = nil
			} else {
				// Если коэффициент не обабатывает руну, то строка не валидна.
				return "", errors.New("Invalid format")
			}
		}
	}

	// Сборка строки и возврат.
	return builder.String(), nil
}
