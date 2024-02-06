package main

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	command := "ls -l|ls -a|ls"

	for _, cmd := range getTokens(command) {
		var this *exec.Cmd

		if len(cmd) == 1 {
			this = exec.Command(cmd[0])
		} else {
			this = exec.Command(cmd[0], cmd[1:]...)
		}

		output, err := this.Output()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(output))
	}

}

// Обрабатывает как конвеер команд, так и единичную команду, отделяя имя команды от аргументов.
func getTokens(pipe string) [][]string {
	cmds := strings.Split(pipe, "|")
	tokens := make([][]string, len(cmds))

	for i := 0; i < len(cmds); i++ {
		token := strings.Trim(cmds[i], " ")
		parts := strings.Split(token, " ")
		tokens[i] = parts
	}

	return tokens
}
