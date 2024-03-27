package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	expectedResult := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}

	result := findAnagramSets(words)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Результаты не совпадают.\nОжидаемый результат: %v\nФактический результат: %v", expectedResult, result)
	}
}

func TestSortString(t *testing.T) {
	input := "пятак"
	expectedResult := "акптя"

	result := sortString(input)

	if result != expectedResult {
		t.Errorf("Результаты не совпадают.\nОжидаемый результат: %v\nФактический результат: %v", expectedResult, result)
	}
}
