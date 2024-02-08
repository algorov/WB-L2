package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Объявляет переменные для парсинга флагов.
	var (
		f int
		d string
		s bool
	)

	// Определение того, что парсить.
	flag.IntVar(&f, "f", 0, "Выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "Только строки с разделителем")

	// Парсит.
	flag.Parse()

	// Слайс с обработанными данными.
	data := make([][]string, 0, 10)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Чтобы выйти: 'quit'")

	for {
		// Сканирует stdin.
		scanner.Scan()
		line := scanner.Text()

		// Если пользователь написал стоп-слово, то прекращает считывание.
		if line == "quit" {
			break
		}

		// В остальных случаях кроме, если в строке нет разделителя, но указан флаг, добавляет строку для последующего анализа.
		if !(s && !strings.Contains(line, d)) {
			data = append(data, strings.Split(line, d))
		}

	}

	// Если пользователь не указал флаг или ввёл отрицательное значение, то выводит все введенное пользователем данные c разделителем.
	if f <= 0 {
		for _, line := range data {
			for _, word := range line {
				fmt.Print(word + d)
			}

			fmt.Println()
		}
	} else {
		// Если пользователь ввёл валиднкую колонку.
		for _, line := range data {
			// Проверка валидности индекса
			if f < len(line) {
				fmt.Println(line[f-1])
			}
		}
	}
}
