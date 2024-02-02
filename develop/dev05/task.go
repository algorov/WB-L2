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
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var v bool
	var a int

	substring := "пустынном"

	flag.BoolVar(&v, "v", false, "Вместо совпадения, исключать")
	flag.IntVar(&a, "A", 0, "Печатать +N строк после совпадения")

	flag.Parse()

	data, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Инициализирует правило валидирования строк по умолчанию.
	var validate func(word, substring string) bool = func(word, substring string) bool {
		return word == substring
	}

	// Если активен флаг V, то переопределяем правило отбора.
	if v {
		validate = func(word, substring string) bool {
			return word != substring
		}
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if validate(data[i][j], substring) {
				fmt.Println(data[i][j])
			}
		}
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
