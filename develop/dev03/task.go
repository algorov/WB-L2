package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки [+]
-n — сортировать по числовому значению [+]
-r — сортировать в обратном порядке [+]
-u — не выводить повторяющиеся строки [ ]

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца [+]
-b — игнорировать хвостовые пробелы [ ]
-c — проверять отсортированы ли данные [+]
-h — сортировать по числовому значению с учётом суффиксов [+]

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Переменные для хранения значения переданных ключей.
	var k int
	var n, r, u, M, b, c, h bool

	// Объявления ключей.
	// Поддержка базовых ключей.
	flag.IntVar(&k, "k", 0, "Указание колонки для сортировки")

	flag.BoolVar(&n, "n", false, "Cортировка по числовому значению")
	flag.BoolVar(&r, "r", false, "Cортировка в обратном порядке")
	flag.BoolVar(&u, "u", false, "Исключение повторяющихся строк")

	// Поддержка дополнительных ключей.
	flag.BoolVar(&M, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&b, "b", false, "Игнорирование хвостовых пробелов")
	flag.BoolVar(&c, "c", false, "Проверка на отсортированность данных")
	flag.BoolVar(&h, "h", false, "Сортировка по числовому значению с учетом суффиксов")

	// Получение ключей.
	flag.Parse()

	// Получение входного и выходного файлов.
	// Проверка на корректность.
	input, output := flag.Arg(0), flag.Arg(1)
	if input == "" || output == "" {
		fmt.Println("Обозначьте, мистер, входной и выходной файлы")
		return
	}

	// Если пользователь указал номер колонки меньше нули или ноль, то колонка выбирается по-умолчанию - 0.
	// Иначе номер колонки декрементируется, поскольку индекс для слайсов начинается с нуля.
	if k <= 0 {
		k = 0
	} else {
		k--
	}

	// Чтение из входного файла.
	data, err := readFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	var compareFunc func(i, j int) bool

	// Составляет функцию для сравнивания значений исходя из переданных флагов.
	switch true {
	// Если сортировка по числовому значению.
	case n:
		compareFunc = func(i, j int) bool {
			// Парсит строку и пробует перевести в числовое значение.
			a, _ := strconv.ParseFloat(getElement(data, i, k), 64)
			b, _ := strconv.ParseFloat(getElement(data, j, k), 64)

			// Если присутствует флаг «r».
			if r {
				return a > b
			}

			return a < b
		}
	case h:
		compareFunc = func(i, j int) bool {
			// Извлекает из набора отдельные элементы.
			a, b := getElement(data, i, k), getElement(data, j, k)

			// Пытается найти совпадение по регулярному выражению.
			valA, errA := extractNumericValue(a)
			valB, errB := extractNumericValue(b)

			// Проверка на неудачу при поиске.
			if errA != nil || errB != nil {
				// Обработка ошибок, например, если извлечение числового значения не удалось
				// Если присутствует флаг «r».
				if r {
					return a > b
				}

				return a < b
			}

			// Если присутствует флаг «r».
			if r {
				return valA > valB
			}

			return valA < valB
		}
	// Сортировка по названию месяца.
	case M:
		compareFunc = func(i, j int) bool {
			// Если присутствует флаг «r».
			if r {
				return parseMonth(getElement(data, j, k)).Before(parseMonth(getElement(data, i, k)))
			}

			return parseMonth(getElement(data, i, k)).Before(parseMonth(getElement(data, j, k)))
		}

	// Если никакого режима сортировки нет, то по умолчанию пойдет обычная сортировка
	default:
		compareFunc = func(i, j int) bool {
			// Если присутствует флаг «r».
			if r {
				return getElement(data, i, k) > getElement(data, j, k)
			}

			return getElement(data, i, k) < getElement(data, j, k)
		}
	}

	// Проверка: сортированы ли наши данные?
	if c {
		if sort.SliceIsSorted(data, compareFunc) {
			fmt.Println("Данные отсортированы!")
		} else {
			fmt.Println("Данные не отсортированы!")
		}
	}

	// Сортировка строк по колонке на основе полученного правила для сортировки.
	sort.Slice(data, compareFunc)

	// Запись в выходной файл.
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
func getElement(data [][]string, i, k int) string {
	// Если индекс допустим в слайсе, то возвращает элемент.
	if k >= 0 && k < len(data[i]) {
		return data[i][k]
	}

	// Иначе возвращает пустую строку.
	return ""
}

// Парсит название месяца из строки.
func parseMonth(str string) time.Time {
	if t, err := time.Parse("January", str); err == nil {
		return t
	}

	if t, err := time.Parse("Jan", str); err == nil {
		return t
	}

	if t, err := time.Parse("1", str); err == nil {
		return t
	}

	if t, err := time.Parse("01", str); err == nil {
		return t
	}

	// Возвращает значение по умолчанию.
	return time.Time{}
}

// Извлекает числовое значение
func extractNumericValue(s string) (int, error) {
	// Пытается найти совпадение по регулярному выражению.
	re := regexp.MustCompile(`(\d+)[KkMm]*`)
	match := re.FindStringSubmatch(s)

	// Если есть совпадение.
	if len(match) == 2 {
		return strconv.Atoi(match[1])
	}

	// Если нет совпадения.
	return 0, fmt.Errorf("Не удалось извлечь числовое значение: %s", s)
}
