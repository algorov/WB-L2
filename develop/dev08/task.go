package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то) [+]
- pwd - показать путь до текущего каталога [+]
- echo <args> - вывод аргумента в STDOUT [+]
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

// Харнит текущую директорию
var curDir string

func main() {
	// Получает текущую рабочую директорию.
	updatePwd()

	for {
		// Стандартный вывод в терминале.
		fmt.Printf("\033[1;32;40m%s:\033[0m ", curDir)

		// Чтобы считывать с stdin.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		// Разбиение конвеера на отдельные команды
		tokens := getTokens(scanner.Text())
		for _, cmd := range tokens {
			switch cmd[0] {
			case "cd":
				if len(cmd) > 1 {
					// Меняет рабочую директорию.
					if err := os.Chdir(cmd[1]); err != nil {
						fmt.Printf("\033[1;31;40m%s\033[0m\n", err)
						continue
					}

					// Обновление пути до текущей рабочей директории.
					updatePwd()
				}
			default:
				// Если получена пустая строка (то есть, если пользователь просто enter нажал).
				if cmd[0] == "" {
					continue
				}

				// Выполнение команды.
				answer, err := execCmd(cmd...)
				if err != nil {
					fmt.Printf("\033[1;31;40m%s\033[0m\n", err)
					continue
				}

				// Вывод результата команды.
				fmt.Println(answer)
			}
		}
	}

}

// Обновляет путь до текущей директории.
func updatePwd() {
	// Вычисляет текущую рабочую директорию.
	dir, _ := os.Getwd()

	// Присваивает переменной.
	curDir = dir
}

// Выполняет команду.
func execCmd(cmd ...string) (string, error) {
	var this *exec.Cmd

	// Если команда без аргументов.
	if len(cmd) == 1 {
		this = exec.Command(cmd[0])
	} else {
		this = exec.Command(cmd[0], cmd[1:]...)
	}

	// Получение ответа.
	output, err := this.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Обрабатывает как конвеер команд, так и единичную команду, отделяя команду от аргументов.
func getTokens(pipe string) [][]string {
	// Разделяет (если есть) конвеер команд.
	cmds := strings.Split(pipe, "|")

	// Отдельные команды с аргументами будут храниться.
	tokens := make([][]string, len(cmds))

	// Убирает лишние пробелы и структурирует команды с аргументами.
	for i := 0; i < len(cmds); i++ {
		token := strings.Trim(cmds[i], " ")
		parts := strings.Split(token, " ")
		tokens[i] = parts
	}

	return tokens
}
