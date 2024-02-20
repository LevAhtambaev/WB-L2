package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var ErrIncorrectString = errors.New("incorrect string")

// Unpack принимает строку str и возвращает распакованную строку
func Unpack(str string) (string, error) {
	runes := []rune(str)
	var letter rune
	var counter = 1
	builder := strings.Builder{}

	// Итерация по всем рунам в строке
	for i := 0; i < len(runes); {
		if !isLetter(runes[i]) {
			return "", ErrIncorrectString
		}
		letter = runes[i]
		i++

		// Подсчет количества повторений символа
		if i >= len(runes) || isLetter(runes[i]) {
			counter = 1
		} else {
			counter = 0
			for i < len(runes) && !isLetter(runes[i]) {
				counter = counter*10 + int(runes[i]-'0')
				i++
			}
		}

		// Добавляем повторения символа на выходную строку
		for j := 0; j < counter; j++ {
			builder.Write([]byte{byte(letter)})
		}
	}
	return builder.String(), nil
}

func isLetter(r rune) bool {
	return r < '0' || r > '9'
}

func main() {
	fmt.Println(Unpack("a4bc2d5e")) // Вывод: "aaaabccddddde" nil
	fmt.Println(Unpack("abcd"))     // Вывод: "abcd" nil
	fmt.Println(Unpack("45"))       // Вывод: "" incorrect string
	fmt.Println(Unpack(""))         // Вывод: "" nil

}
