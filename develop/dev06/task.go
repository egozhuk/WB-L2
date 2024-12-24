package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type cutConfig struct {
	fields    []int  // Номера колонок для вывода
	delimiter string // Разделитель
	separated bool   // Только строки с разделителем
}

// cut выполняет обработку входных строк на основе конфигурации
func cut(input []string, config cutConfig) []string {
	var result []string

	for _, line := range input {
		// Проверяем наличие разделителя
		if !strings.Contains(line, config.delimiter) && config.separated {
			continue
		}

		// Разбиваем строку на колонки
		columns := strings.Split(line, config.delimiter)

		// Выбираем только запрошенные поля
		var selectedColumns []string
		for _, field := range config.fields {
			if field > 0 && field <= len(columns) {
				selectedColumns = append(selectedColumns, columns[field-1])
			}
		}

		// Добавляем результат
		result = append(result, strings.Join(selectedColumns, config.delimiter))
	}

	return result
}

func parseFields(fields string) ([]int, error) {
	var result []int
	for _, f := range strings.Split(fields, ",") {
		var field int
		_, err := fmt.Sscanf(f, "%d", &field)
		if err != nil {
			return nil, fmt.Errorf("invalid field: %v", f)
		}
		result = append(result, field)
	}
	return result, nil
}

func main() {
	// Определяем флаги
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Set the delimiter")
	separated := flag.Bool("s", false, "Only process lines with delimiter")
	flag.Parse()

	// Проверяем аргументы
	if *fields == "" {
		fmt.Fprintln(os.Stderr, "Error: -f flag is required")
		os.Exit(1)
	}

	// Парсим номера полей
	parsedFields, err := parseFields(*fields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing fields: %v\n", err)
		os.Exit(1)
	}

	// Читаем строки из STDIN
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Конфигурация
	config := cutConfig{
		fields:    parsedFields,
		delimiter: *delimiter,
		separated: *separated,
	}

	// Выполняем cut
	output := cut(input, config)

	// Печатаем результат
	for _, line := range output {
		fmt.Println(line)
	}
}
