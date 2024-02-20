package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	goCmdArgs := [][]string{
		{"run", "task.go", "-f=2,5"},
		{"run", "task.go", "-f=2,5", "-d=;"},
		{"run", "task.go", "-f=2,5", "-d=;", "-s"},
	}

	var stdout bytes.Buffer

	input := []string{
		"1 apple banana cat dog car window",
		"2; apple; banana; cat; dog; car; window:",
		"3;apple;banana;cat;dog;car;window;",
	}

	expected := []string{
		"apple dog\n",
		" apple; dog\n",
		"apple;dog\n",
	}

	for i := range goCmdArgs {
		cmd := exec.Command("go", goCmdArgs[i]...)
		cmd.Stdin = strings.NewReader(input[i])
		stdout.Reset()
		cmd.Stdout = &stdout

		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}

		if res := stdout.String(); res != expected[i] {
			t.Errorf("actual: %v, excepted: %v", res, expected[i])
		}
	}
}
