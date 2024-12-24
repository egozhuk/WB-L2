package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type grepConfig struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func grep(input []string, config grepConfig) []string {
	var result []string
	matchLines := map[int]bool{}

	// Преобразуем паттерн, если включен флаг ignoreCase
	pattern := config.pattern
	if config.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	for i, line := range input {
		if config.ignoreCase {
			line = strings.ToLower(line)
		}

		// Проверяем, есть ли совпадение
		match := false
		if config.fixed {
			match = line == pattern
		} else {
			match = strings.Contains(line, pattern)
		}

		// Инвертируем результат, если установлен флаг invert
		if config.invert {
			match = !match
		}

		// Если совпадение найдено
		if match {
			matchLines[i] = true
		}
	}

	// Если флаг count установлен, возвращаем количество совпадений
	if config.count {
		result = append(result, strconv.Itoa(len(matchLines)))
		return result
	}

	// Обрабатываем контекст (до, после, вокруг)
	contextLines := map[int]bool{}
	for i := range matchLines {
		start := i - config.before - config.context
		end := i + config.after + config.context
		for j := start; j <= end; j++ {
			if j >= 0 && j < len(input) {
				contextLines[j] = true
			}
		}
	}

	// Сортируем индексы строк для вывода в правильном порядке
	var sortedIndexes []int
	for i := range contextLines {
		sortedIndexes = append(sortedIndexes, i)
	}
	sort.Ints(sortedIndexes)

	// Формируем результат
	for _, i := range sortedIndexes {
		if config.lineNum {
			result = append(result, fmt.Sprintf("%d:%s", i+1, input[i]))
		} else {
			result = append(result, input[i])
		}
	}

	return result
}

func main() {
	// Определяем флаги
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines before match")
	context := flag.Int("C", 0, "Print ±N lines around match")
	count := flag.Bool("c", false, "Print the count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNum := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	// Оставшиеся аргументы — это паттерн и имя файла
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN FILE")
		os.Exit(1)
	}
	pattern := args[0]
	fileName := args[1]

	// Читаем файл
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Обрабатываем файл с учетом конфигурации
	config := grepConfig{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		pattern:    pattern,
	}
	result := grep(lines, config)

	// Выводим результат
	for _, line := range result {
		fmt.Println(line)
	}
}
