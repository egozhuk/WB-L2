package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Разбор аргументов
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet --timeout=<duration> host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Устанавливаем соединение
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	// Канал для завершения программы
	done := make(chan struct{})

	// Чтение из сокета и вывод в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Connection closed by server: %v\n", err)
		}
		close(done)
	}()

	// Запись из STDIN в сокет
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to server: %v\n", err)
		}
		conn.Close() // Закрываем соединение при завершении ввода
	}()

	// Ждем завершения
	<-done
	fmt.Println("\nDisconnected")
}
