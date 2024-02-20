package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		excepted string
		err      error
	}{
		{name: "test 1", input: "a4bc2d5e", excepted: "aaaabccddddde", err: nil},
		{name: "test 2", input: "abcd", excepted: "abcd", err: nil},
		{name: "test 3", input: "45", excepted: "", err: ErrIncorrectString},
		{name: "test 4", input: "", excepted: "", err: nil},
		{name: "test 5", input: "hello10", excepted: "helloooooooooo", err: nil},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			actual, actualErr := Unpack(e.input)
			if actual != e.excepted {
				t.Errorf("actual: %v, excepted: %v", actual, e.excepted)
			}
			if !errors.Is(actualErr, e.err) {
				t.Errorf("actual: %v, excepted: %v", actualErr, e.err)
			}
		})
	}
}
