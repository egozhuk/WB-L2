package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	// Создаём временную директорию для теста
	tempDir, err := ioutil.TempDir("", "testDownload")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Очищаем после теста

	// Создаём тестовый HTTP сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body>Hello, world!</body></html>"))
	}))
	defer ts.Close()

	// Изменяем функцию downloadPage, чтобы она принимала путь к директории
	err = downloadPageToDir(ts.URL, tempDir)
	if err != nil {
		t.Errorf("downloadPage returned an error: %v", err)
	}

	// Проверяем, что файл был создан и содержит правильные данные
	filename := filepath.Join(tempDir, "index.html")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(data) != "<html><body>Hello, world!</body></html>" {
		t.Errorf("File contents did not match expected: got %s", data)
	}
}

// Изменённая функция downloadPage, принимающая директорию для сохранения файла
func downloadPageToDir(pageURL, dir string) error {
	resp, err := http.Get(pageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := filepath.Join(dir, "index.html")

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
