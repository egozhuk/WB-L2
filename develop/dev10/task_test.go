package main

import (
	"net"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	// Создаем тестовый TCP сервер
	ln, err := net.Listen("tcp", "127.0.0.1:0") // 0 означает автоматический выбор свободного порта
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		// Эхо-сервер: отправляем обратно полученные данные
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return // Закрываем соединение при ошибке чтения
			}
			_, err = conn.Write(buf[:n])
			if err != nil {
				return // Закрываем соединение при ошибке записи
			}
		}
	}()

	// Подключаемся к тестовому серверу
	address := ln.Addr().String()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Проверяем отправку и получение данных
	testData := "Hello, telnet!"
	_, err = conn.Write([]byte(testData))
	if err != nil {
		t.Fatal(err)
	}

	buffer := make([]byte, len(testData))
	_, err = conn.Read(buffer)
	if err != nil {
		t.Fatal(err)
	}

	if string(buffer) != testData {
		t.Errorf("Expected %q, got %q", testData, string(buffer))
	}
}

func TestTelnetClientTimeout(t *testing.T) {
	// Проверяем таймаут подключения
	_, err := net.DialTimeout("tcp", "192.0.2.1:80", 1*time.Millisecond) // 192.0.2.1 - адрес для тестирования
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}
