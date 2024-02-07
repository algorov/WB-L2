package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Чтобы считывать с stdin.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the URL: ")
	scanner.Scan()

	// Скачивание содержимого из переданного URL.
	if err := Download(scanner.Text()); err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

// Сохраняет содержимое, полученное от сервера по переданному адресу.
func Download(url string) error {
	// Запрос.
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// Нужен для имени файла
	temp := strings.Split(url, "/")

	// Формирование имени файла
	filename := temp[len(temp)-1] + ".html"

	// Создает файл с сформированным именем.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копирует в созданный файл ответ от сервера.
	io.Copy(file, response.Body)

	return nil
}
