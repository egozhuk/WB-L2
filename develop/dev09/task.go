package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run wget.go <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	err := downloadPage(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download complete.")
}

func downloadPage(url string) error {
	// Отправляем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Определяем имя файла
	filename := filepath.Base(url)
	if filename == "" || filename == "/" {
		filename = "index.html"
	}

	// Создаем файл для сохранения
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Сохраняем содержимое ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
