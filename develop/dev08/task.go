package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Welcome to Simple Shell! Type \\quit to exit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Печатаем приглашение
		fmt.Print("> ")

		// Читаем ввод пользователя
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Удаляем лишние пробелы и символы новой строки
		input = strings.TrimSpace(input)

		// Команда выхода
		if input == "\\quit" {
			fmt.Println("Exiting Simple Shell.")
			break
		}

		// Разбиваем команды по пайпам
		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			handlePipeline(commands)
		} else {
			handleCommand(input)
		}
	}
}

// handleCommand обрабатывает одиночные команды
func handleCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Println("Usage: cd <directory>")
			return
		}
		if err := os.Chdir(args[1]); err != nil {
			fmt.Println("Error changing directory:", err)
		}
	case "pwd":
		if dir, err := os.Getwd(); err == nil {
			fmt.Println(dir)
		} else {
			fmt.Println("Error getting current directory:", err)
		}
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			fmt.Println("Usage: kill <pid>")
			return
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid PID:", args[1])
			return
		}
		if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
			fmt.Println("Error killing process:", err)
		}
	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error executing ps:", err)
		}
	default:
		execCommand(args)
	}
}

// handlePipeline обрабатывает команды с пайпами
func handlePipeline(commands []string) {
	var prevCmd *exec.Cmd

	for i, command := range commands {
		args := strings.Fields(strings.TrimSpace(command))
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)

		if i == 0 {
			// Первая команда читает из стандартного ввода
			cmd.Stdin = os.Stdin
		} else {
			// Все остальные команды читают из канала предыдущей
			cmd.Stdin, _ = prevCmd.StdoutPipe()
		}

		// Если последняя команда, то пишет в стандартный вывод
		if i == len(commands)-1 {
			cmd.Stdout = os.Stdout
		} else {
			cmd.Stdout = nil
		}

		// Запускаем команду
		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		// Ждем завершения предыдущей команды
		if prevCmd != nil {
			_ = prevCmd.Wait()
		}

		prevCmd = cmd
	}

	// Ждем завершения последней команды
	if prevCmd != nil {
		_ = prevCmd.Wait()
	}
}

// execCommand выполняет внешнюю команду
func execCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command '%s': %v\n", args[0], err)
	}
}
