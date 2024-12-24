package main

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)
	uniqueWords := make(map[string]bool)

	for _, word := range words {
		// Приводим слово к нижнему регистру
		wordLower := strings.ToLower(word)
		if uniqueWords[wordLower] {
			continue
		}
		uniqueWords[wordLower] = true

		// Сортируем буквы в слове
		key := sortString(wordLower)

		// Добавляем слово в соответствующую группу
		anagramGroups[key] = append(anagramGroups[key], wordLower)
	}

	// Формируем результат
	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) > 1 {
			sort.Strings(group) // Сортируем группу по алфавиту
			result[group[0]] = group
		}
	}

	return result
}

// sortString сортирует буквы в строке
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "тяпка", "ПЯТАК"}
	anagrams := findAnagrams(words)

	for key, group := range anagrams {
		println("Key:", key, "Group:", strings.Join(group, ", "))
	}
}
