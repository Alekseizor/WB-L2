package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagramSets(words *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)
	uniqueWords := make(map[string]bool)
	for _, word := range *words {
		//для получения множеств проверяем на уникальность
		if uniqueWords[word] {
			continue
		}
		uniqueWords[word] = true
		// Приводим слово к нижнему регистру и сортируем его буквы
		word = strings.ToLower(word)
		sortedWord := sortString(word)

		// Добавляем отсортированное слово в множество анаграмм
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	result := make(map[string][]string)

	for _, wordsResult := range anagramSets {
		if len(wordsResult) > 1 {
			firstWord := wordsResult[0]
			sort.Strings(wordsResult)
			result[firstWord] = wordsResult
		}
	}

	return &result
}

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func main() {
	words := []string{"пятка", "пятак", "тяпка", "тяпка", "слиток", "листок", "столик"}
	anagramSets := FindAnagramSets(&words)

	for key, words := range *anagramSets {
		fmt.Println(key, words)
	}
}
