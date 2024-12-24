package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Basic anagrams",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "No anagrams",
			input:    []string{"один", "два", "три"},
			expected: map[string][]string{},
		},
		{
			name:  "Mixed case",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "тяпка", "ПЯТАК"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "Duplicates in input",
			input: []string{"пятак", "пятка", "тяпка", "пятак", "пятка"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "Single word",
			input:    []string{"один"},
			expected: map[string][]string{},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findAnagrams(test.input)

			// Проверяем количество ключей в результате
			if len(result) != len(test.expected) {
				t.Errorf("Expected %d groups, got %d", len(test.expected), len(result))
			}

			// Проверяем каждую группу анаграмм
			for key, expectedGroup := range test.expected {
				group, exists := result[key]
				if !exists {
					t.Errorf("Expected group with key %q not found", key)
					continue
				}

				// Сравниваем отсортированные группы
				if !reflect.DeepEqual(group, expectedGroup) {
					t.Errorf("For key %q, expected group: %v, got: %v", key, expectedGroup, group)
				}
			}
		})
	}
}
