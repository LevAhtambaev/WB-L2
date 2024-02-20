package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// sortLetters сортирует буквы в слове
func sortLetters(s string) string {
	r := []rune(s)
	compare := func(i, j int) bool {
		r1 := r[i]
		r2 := r[j]
		res := strings.Compare(string(r1), string(r2))
		return res < 0
	}

	sort.SliceStable(r, compare)

	return string(r)
}

// makeAnagrams возвращает анаграммы из заданного набора слов
func makeAnagrams(input []string) map[string][]string {
	words := make(map[string][]string)
	anagrams := make(map[string][]string)

	for _, v := range input {
		v = strings.ToLower(v)
		sortedWord := sortLetters(v)
		if _, ok := words[sortedWord]; ok {
			words[sortedWord] = append(words[sortedWord], v)
		} else {
			words[sortedWord] = []string{v}
		}
	}

	for _, v := range words {
		if len(v) > 1 {
			anagrams[v[0]] = v[1:]
		}
	}

	for _, v := range anagrams {
		sort.Strings(v)
	}

	return anagrams
}

func main() {
	input := []string{"ток", "пятак", "тяпка", "кот", "столик", "листок", "пятка", "слиток"}

	anagrams := makeAnagrams(input)

	for key, v := range anagrams {
		fmt.Printf("Key: %v, value: %v\n", key, v)
	}
}
