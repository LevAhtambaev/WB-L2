package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// getColumn извлекает колонку из строки
func getColumn(line string, column int) string {
	fields := strings.Fields(line)
	// Проверка на диапазон слов в строке
	if column > 0 && column <= len(fields) {
		return fields[column-1]
	}

	return line
}

// stripDuplicates удаляет повторяющиеся строки
func stripDuplicates(lines []string) []string {
	unique := make(map[string]bool)
	var result []string

	for _, v := range lines {
		if _, exist := unique[v]; !exist {
			unique[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	// Флаги для запуска
	k := flag.Int("k", 0, "Указание колонки для сортировки")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	// Смотрим файл для сортировки
	fileUnsort, err := os.Open("for_sort.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fileUnsort.Close()

	var lines []string
	scanner := bufio.NewScanner(fileUnsort)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

	// Сравниваем строки для сортировки
	compare := func(i, j int) bool {
		s1 := getColumn(lines[i], *k)
		s2 := getColumn(lines[j], *k)

		// -n - сортировка чисел
		if *n {
			num1, err1 := strconv.Atoi(s1)
			num2, err2 := strconv.Atoi(s2)
			if err1 == nil && err2 == nil {
				switch {
				case *r && num1 < num2:
					return false
				case *r && num1 > num2:
					return true
				case !*r && num1 < num2:
					return true
				case !*r && num1 > num2:
					return false
				}
			}
		}

		result := strings.Compare(s1, s2)

		// -r - обратный порядок
		if *r {
			return result > 0
		}

		return result < 0
	}

	// -u - удаляем повторы
	if *u {
		lines = stripDuplicates(lines)
	}

	sort.SliceStable(lines, compare)

	// Записываем отсортированный файл
	fileSort, err := os.Create("after_sort.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fileSort.Close()

	for idx, value := range lines {
		fileSort.WriteString(value)
		if idx != len(lines)-1 {
			fileSort.WriteString("\n")
		}
	}
}
