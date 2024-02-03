package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения [+]
-B - "before" печатать +N строк до совпадения [+]
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк) [+]
-i - "ignore-case" (игнорировать регистр) [+]
-v - "invert" (вместо совпадения, исключать) [+]
-F - "fixed", точное совпадение со строкой, не паттерн [+]
-n - "line num", печатать номер строки [+]

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var a, b, context int
	var c, i, v, f, n bool

	substring := "пустын"

	flag.IntVar(&a, "A", 0, "Печатать +N строк после совпадения")
	flag.IntVar(&b, "B", 0, "Печатать -N строк до совпадения")
	flag.IntVar(&context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")

	flag.BoolVar(&c, "c", false, "Количество строк")
	flag.BoolVar(&i, "i", false, "Игнорировать регистр")
	flag.BoolVar(&v, "v", false, "Вместо совпадения, исключать")
	flag.BoolVar(&f, "f", false, "Точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "Печатать номер строки")

	flag.Parse()

	data, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Инициализирует правило подготовки данных для последующего анализа.
	var prepare func(word, substring string) (string, string) = func(word, substring string) (string, string) {
		return word, substring
	}

	// Инициализирует правило поиска подстроки в строке.
	var isContain func(word, substring string) bool = func(word, substring string) bool {
		return strings.Contains(word, substring)
	}

	// Инициализирует правило отбора результата по умолчанию.
	var pass func(isContain bool) bool = func(isContain bool) bool {
		return isContain
	}

	// Если активен флаг i, то переопределяем правило подготовки данных.
	if i {
		prepare = func(word, substring string) (string, string) {
			return strings.ToLower(word), strings.ToLower(substring)
		}
	}

	// Если активен флаг i, то переопределяем правило поиска.
	if f {
		isContain = func(word, substring string) bool {
			return word == substring
		}
	}

	// Если активен флаг V, то переопределяем правило отбора.
	if v {
		pass = func(isContain bool) bool {
			return !isContain
		}
	}

	// Место, где хранятся совпадения на основе полученных правил отбора.
	result := make([]int, 0, len(data))

	// Применяет правило на данных из входного файла.
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if pass(isContain(prepare(data[i][j], substring))) {
				result = append(result, i)
			}
		}
	}

	// Обрабротка флага c.
	if c {
		fmt.Println("Line count:", len(result))
	}

	// Если никаких совпадений нет.
	if len(result) == 0 {
		fmt.Println("Совпадений нет!")

		return
	}

	// Нормализация значение флага A.
	if a < 0 {
		a = 0
	}

	// Нормализация значение флага B.
	if b < 0 {
		b = 0
	}

	// Нормализация значение флага С.
	if context < 0 {
		context = 0
	}

	switch true {
	// Обработка флага C.
	case context > 0:
		a, b = context, context
		printBefore(result[0], b, data)
		printAfter(result[0], a, data)
	// Обработка флага A.
	case a > 0:
		printAfter(result[0], a, data)
	// Обработка флага B.
	case b > 0:
		printBefore(result[0], b, data)
	// Обработка флага n.
	case n:
		for i, _ := range result {
			fmt.Printf("[%d]: %s\n", result[i]+1, data[result[i]])
		}
	// Печатает первое совпадение, если, конечно, есть.
	default:
		fmt.Printf("[%d]: %s\n", result[0]+1, data[result[0]])
	}
}

// Чтение данных из файла.
func readFile(path string) (data [][]string, err error) {
	// Получение файлового дескриптора.
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Сканнер для чтения строки.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Добавление в слайс разделенной пробелом строки.
		data = append(data, strings.Split(scanner.Text(), " "))
	}

	// Возврат результата.
	return data, nil
}

// Выводит N записей до совпадения.
func printBefore(current, count int, data [][]string) {
	side := current - count
	// Нормализация границы в пределах допустимых значений индексов
	if side < 0 {
		side = 0
	}

	for index := side; index < current+1; index++ {
		fmt.Printf("[%d]: %s\n", index+1, data[index])
	}
}

// Выводит N записей после совпадения.
func printAfter(current, count int, data [][]string) {
	side := current + count
	// Нормализация границы в пределах допустимых значений индексов
	if side > len(data) {
		side = len(data)
	}

	for index := current; index < side; index++ {
		fmt.Printf("[%d]: %s\n", index+1, data[index])
	}
}
