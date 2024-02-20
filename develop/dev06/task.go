package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// getFields возвращает указанные поля
func getFields(fields []string, fieldsList string) []string {
	selectedFields := make([]string, 0)
	fieldsIdx := getIndexes(fieldsList)

	for _, idx := range fieldsIdx {
		if idx > 0 && idx <= len(fields) {
			selectedFields = append(selectedFields, fields[idx-1])
		} else {
			selectedFields = append(selectedFields, "")
		}
	}

	return selectedFields
}

// getIndexes возвращает номера полей
func getIndexes(fieldsList string) []int {
	fieldsIdx := make([]int, 0)
	fields := strings.Split(fieldsList, ",")

	for _, field := range fields {
		idx := 0
		if n, err := fmt.Sscanf(field, "%d", &idx); err == nil && n > 0 {
			fieldsIdx = append(fieldsIdx, idx)
		}
	}

	return fieldsIdx
}

func main() {
	f := flag.String("f", "", "выбрать поля (колонки)")
	d := flag.String("d", " ", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *s || strings.Contains(line, *d) {
			fields := strings.Split(line, *d)

			if *f != "" {
				selectedFields := getFields(fields, *f)
				fmt.Println(strings.Join(selectedFields, *d))
			} else {
				fmt.Println(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
