package main

import (
	"reflect"
	"testing"
)

func TestSortLetters(t *testing.T) {
	input := []string{"память", "ток", "лошадь", "орел"}
	expected := []string{"амптья", "кот", "адлошь", "елор"}

	for i := 0; i < len(input); i++ {
		res := sortLetters(input[i])
		if expected[i] != res {
			t.Errorf("actual: %v, excepted: %v", res, expected)
		}
	}
}

func TestMakeAnagrams(t *testing.T) {
	input := []string{"ток", "пятак", "тяпка", "кот", "КтО", "столик", "лиСток", "пятка", "слиток", "молоток"}
	expected := map[string][]string{
		"ток":    {"кот", "кто"},
		"пятак":  {"пятка", "тяпка"},
		"столик": {"листок", "слиток"},
	}

	res := makeAnagrams(input)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("actual: %v, excepted: %v", res, expected)
	}
}
