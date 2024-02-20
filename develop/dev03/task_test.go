package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestGetColumn(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		column   int
	}{
		{
			input:    "ford 2 12",
			expected: "2",
			column:   2,
		},
		{
			input:    "toyota 1",
			expected: "toyota",
			column:   1,
		},
	}
	for _, test := range tests {
		res := getColumn(test.input, test.column)
		if res != test.expected {
			t.Errorf("actual: %v, excepted: %v", res, test.expected)
		}
	}
}

func TestStripDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{name: "test1.txt", input: []string{"test", "test1.txt", "test2.txt"}, expected: []string{"test", "test1.txt", "test2.txt"}},
		{name: "test2.txt", input: []string{"test", "test", "test2.txt", "test2.txt", "test2.txt"}, expected: []string{"test", "test2.txt"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := stripDuplicates(test.input)
			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("actual: %v, excepted: %v", res, test.expected)
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		number int
		args   []string
	}{
		{number: 1, args: []string{"run", "task.go", "-r"}},
		{number: 2, args: []string{"run", "task.go", "-k=1"}},
		{number: 3, args: []string{"run", "task.go", "-u"}},
	}

	for _, test := range tests {
		cmd := exec.Command("go", test.args...)
		cmd.CombinedOutput()

		name := fmt.Sprintf("test%v.txt", test.number)
		testFile, err := os.Open(name)
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer testFile.Close()

		soretedFIle, err := os.Open("after_sort.txt")
		if err != nil {
			t.Errorf("Ошибка открытия файла: %v", err)
		}
		defer soretedFIle.Close()

		expected, _ := io.ReadAll(testFile)
		result, _ := io.ReadAll(soretedFIle)

		if !bytes.Equal(result, expected) {
			t.Errorf("actual: %v\n excepted: %v", string(result), string(expected))
		}
	}

}
