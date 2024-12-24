package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Unpack выполняет распаковку строки с поддержкой экранирования
func Unpack(input string) (string, error) {
	var (
		builder    strings.Builder // Строитель результата
		escapeMode bool            // Флаг экранирования
		lastRune   rune            // Последний символ
	)

	for _, char := range input {
		// Если это цифра и мы не в режиме экранирования
		if unicode.IsDigit(char) && !escapeMode {
			// Проверяем, есть ли символ для клонирования
			if lastRune == 0 {
				return "", errors.New("некорректная строка: цифра без символа перед ней")
			}
			// Получаем количество повторений
			count, err := strconv.Atoi(string(char))
			if err != nil {
				return "", errors.New("ошибка преобразования числа")
			}
			// Дополняем результат
			builder.WriteString(strings.Repeat(string(lastRune), count-1))
			lastRune = 0
			continue
		}

		// Если встречаем '\', включаем режим экранирования
		if char == '\\' && !escapeMode {
			escapeMode = true
			continue
		}

		// Добавляем текущий символ в результат
		builder.WriteRune(char)
		lastRune = char
		escapeMode = false
	}

	// Если строка заканчивается на '\', это ошибка
	if escapeMode {
		return "", errors.New("некорректная строка: незавершенная escape-последовательность")
	}

	return builder.String(), nil
}

func main() {
	tests := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
	}

	for _, test := range tests {
		result, err := Unpack(test)
		if err != nil {
			fmt.Printf("Input: %q => Error: %v\n", test, err)
		} else {
			fmt.Printf("Input: %q => Output: %q\n", test, result)
		}
	}
}
