package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем буквы
		sortedWord := sortString(strings.ToLower(word))

		// Добавляем слово в соответствующее множество анаграмм
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}

// Вспомогательная функция для сортировки букв в слове
func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
func main() {
	// Пример входных данных
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	anagramSets := findAnagramSets(words)

	// Вывод результатов
	for key, value := range anagramSets {
		fmt.Println(key, ": ", value)
	}
}

/*
В данном примере функция findAnagramSets принимает массив слов words и возвращает мапу anagramSets,
где ключи представляют собой отсортированные буквы слова, а значения - массивы слов-анаграмм.
После того, как мапа сформирована, удаляем множества из одного элемента.
Далее в основной функции main приведен пример использования этой функции, где массив words содержит русские слова.
Результаты поиска множеств анаграмм выводятся на экран.
*/
