package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var column int
	var numeric, reverse, unique, month, ignoreTrailingSpace, checkSorted, humanNumeric bool

	flag.IntVar(&column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&month, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&ignoreTrailingSpace, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&checkSorted, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&humanNumeric, "h", false, "сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	if checkSorted {
		sorted, err := isSorted(os.Stdin, column, numeric, reverse, month, ignoreTrailingSpace, humanNumeric)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		if sorted {
			fmt.Println("Данные отсортированы")
		} else {
			fmt.Println("Данные не отсортированы")
		}
		os.Exit(0)
	}

	lines, err := readLines(os.Stdin)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	sortedLines, err := sortLines(lines, column, numeric, reverse, unique, month, ignoreTrailingSpace, humanNumeric)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	for _, line := range sortedLines {
		fmt.Println(line)
	}
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func sortLines(lines []string, column int, numeric, reverse, unique, month, ignoreTrailingSpace, humanNumeric bool) ([]string, error) {
	sort.SliceStable(lines, func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		if ignoreTrailingSpace {
			line1 = strings.TrimRight(line1, " ")
			line2 = strings.TrimRight(line2, " ")
		}

		fields1 := strings.Fields(line1)
		fields2 := strings.Fields(line2)

		if column > 0 {
			if column > len(fields1) || column > len(fields2) {
				return false
			}
			field1 := fields1[column-1]
			field2 := fields2[column-1]
			if numeric {
				if humanNumeric {
					num1, err := parseHumanNumeric(field1)
					if err != nil {
						return false
					}
					num2, err := parseHumanNumeric(field2)
					if err != nil {
						return false
					}
					return compareNumeric(num1, num2, reverse)
				} else {
					num1, err := strconv.ParseFloat(field1, 64)
					if err != nil {
						return false
					}
					num2, err := strconv.ParseFloat(field2, 64)
					if err != nil {
						return false
					}
					return compareNumeric(num1, num2, reverse)
				}
			} else if month {
				time1, err := parseMonth(field1)
				if err != nil {
					return false
				}
				time2, err := parseMonth(field2)
				if err != nil {
					return false
				}
				return compareTime(time1, time2, reverse)
			} else {
				return compareStrings(field1, field2, reverse)
			}
		} else {
			if numeric {
				if humanNumeric {
					num1, err := parseHumanNumeric(line1)
					if err != nil {
						return false
					}
					num2, err := parseHumanNumeric(line2)
					if err != nil {
						return false
					}
					return compareNumeric(num1, num2, reverse)
				} else {
					num1, err := strconv.ParseFloat(line1, 64)
					if err != nil {
						return false
					}
					num2, err := strconv.ParseFloat(line2, 64)
					if err != nil {
						return false
					}
					return compareNumeric(num1, num2, reverse)
				}
			} else if month {
				time1, err := parseMonth(line1)
				if err != nil {
					return false
				}
				time2, err := parseMonth(line2)
				if err != nil {
					return false
				}
				return compareTime(time1, time2, reverse)
			} else {
				return compareStrings(line1, line2, reverse)
			}
		}
	})

	if unique {
		lines = removeDuplicates(lines)
	}

	return lines, nil
}

func parseHumanNumeric(s string) (float64, error) {
	// TODO: Реализовать парсинг числового значения с учетом суффиксов
	return 0, nil
}

func parseMonth(s string) (time.Time, error) {
	// TODO: Реализовать парсинг названия месяца
	return time.Time{}, nil
}

func compareNumeric(a, b float64, reverse bool) bool {
	if a < b {
		return !reverse
	} else if a > b {
		return reverse
	}
	return false
}

func compareTime(a, b time.Time, reverse bool) bool {
	if a.Before(b) {
		return !reverse
	} else if a.After(b) {
		return reverse
	}
	return false
}

func compareStrings(a, b string, reverse bool) bool {
	if a < b {
		return !reverse
	} else if a > b {
		return reverse
	}
	return false
}

func removeDuplicates(lines []string) []string {
	var uniqueLines []string

	duplicateSet := make(map[string]bool)
	for _, line := range lines {
		if _, ok := duplicateSet[line]; !ok {
			duplicateSet[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}

func isSorted(r io.Reader, column int, numeric, reverse, month, ignoreTrailingSpace, humanNumeric bool) (bool, error) {
	lines, err := readLines(r)
	if err != nil {
		return false, err
	}

	sortedLines, err := sortLines(lines, column, numeric, reverse, false, month, ignoreTrailingSpace, humanNumeric)
	if err != nil {
		return false, err
	}

	return equal(lines, sortedLines), nil
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
