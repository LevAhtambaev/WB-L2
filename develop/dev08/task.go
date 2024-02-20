package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("shell started")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}
		pipeline := strings.Split(input, " | ")

		for _, v := range pipeline {
			args := strings.Fields(v)
			if len(args) == 0 {
				continue
			}

			switch args[0] {
			case "cd":
				if len(args) < 2 {
					home, err := os.UserHomeDir()
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}

					err = os.Chdir(home)
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}
				} else {
					err := os.Chdir(args[1])
					if err != nil {
						fmt.Fprintf(os.Stderr, "cd: %v\n", err)
					}
				}
			case "pwd":
				dir, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
				}

				fmt.Println(dir)
			case "echo":
				fmt.Println(strings.Join(args[1:], " "))
			case "kill":
				if len(args) < 2 {
					fmt.Println("kill: missing argument")
				} else {
					pid, err := strconv.Atoi(args[1])
					if err != nil {
						fmt.Println(err)
					}

					proc, err := os.FindProcess(pid)
					if err != nil {
						fmt.Fprintf(os.Stderr, "kill: %v\n", err)
					}

					err = proc.Kill()
					if err != nil {
						fmt.Fprintf(os.Stderr, "kill: %v\n", err)
					}
				}
			case "ps":
				cmd := exec.Command("ps")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(cmd.Stderr, "ps: %v\n", err)
				}
			default:
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v: %v\n", args[0], err)
				}
			}
		}
	}
}
