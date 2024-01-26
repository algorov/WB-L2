package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var k int
	var n, r, u, M, b, c, h bool

	flag.IntVar(&k, "k", 0, "Указание колонки для сортировки")

	flag.BoolVar(&n, "n", false, "Cортировка по числовому значению")
	flag.BoolVar(&r, "r", false, "Cортировка в обратном порядке")
	flag.BoolVar(&u, "u", false, "Исключение повторяющихся строк")

	flag.BoolVar(&M, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&b, "b", false, "Игнорирование хвостовых пробелов")
	flag.BoolVar(&c, "c", false, "Проверка на отсортированность данных")
	flag.BoolVar(&h, "h", false, "Сортировка по числовому значению с учетом суффиксов")

	flag.Parse()

	input, output := flag.Arg(0), flag.Arg(1)
	if input == "" || output == "" {
		fmt.Println("Обозначьте, мистер, входной и выходной файлы")
		return
	}

	data, err := readFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := writeFile(output, data); err != nil {
		fmt.Println(err)
		return
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

// Запись данных в файл.
func writeFile(fileName string, data [][]string) error {
	// Создает файл с указанным именем и получает файловый дескриптор.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Слайс нужен, чтобы хранить собранные строки.
	lines := make([]string, 0, len(data))

	// В цикле токены собираются в одну строку, которая помещается в слайс строк.
	for i := 0; i < len(data); i++ {
		lines = append(lines, strings.Join(data[i], " "))
	}

	// Запись собранных в одну строку с последующей конвертацией в слайс байтов в файл.
	file.Write([]byte(strings.Join(lines, "\n")))

	return nil
}

// Доступ к определенному элементу.
func getElement(data [][]string, i, j int) string {
	// Если индекс допустим в слайсе, то возвращает элемент.
	if j >= 0 && j < len(data[i]) {
		return data[i][j]
	}

	// Иначе возвращает пустую строку.
	return ""
}
