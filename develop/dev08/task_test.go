package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type StrErr struct {
	input string
	err   error
}
type ArgsOut struct {
	args []string
	out  string
}

func TestShell(t *testing.T) {
	testCases := []StrErr{
		{
			input: "Hello, World!",
			err:   io.EOF,
		},
		{
			input: "ls | sort\nexit\n",
			err:   nil,
		},
		{
			input: "cd develop\nexit\n",
			err:   nil,
		},
		{
			input: "pwd\nexit\n",
			err:   nil,
		},
		{
			input: "echo test\nexit\n",
			err:   nil,
		},
		{
			input: "kill a\nexit\n",
			err:   nil,
		},
		{
			input: "ps\nexit\n",
			err:   nil,
		},
		{
			input: "ls\nexit\n",
			err:   nil,
		},
		{
			input: "test\nexit\n",
			err:   nil,
		},
	}
	for numCase, testCase := range testCases {
		err := Shell(bufio.NewReader(strings.NewReader(testCase.input)))
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCase, err, testCase.err)
		}
	}
}

func TestHandleCd(t *testing.T) {
	testCases := []ArgsOut{
		{
			args: []string{
				"develop",
				"listing",
			},
			out: "Too many arguments\n",
		},
		{
			args: nil,
			out:  "Missing argument for cd command\n",
		},
	}
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Println(err)
		return
	}
	os.Stdout = w
	// Перехватываем вывод в stdOut

	// Читаем данные из перехваченного вывода
	for numCase, testCase := range testCases {
		handleCd(testCase.args)
		buf := make([]byte, 100)
		n, err := r.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		output := string(buf[:n])
		log.Println(output)
		if output != testCase.out {
			t.Errorf("[%d] returned output - %v - different from expected - %v", numCase, output, testCase.out)
		}
	}
	w.Close()
	os.Stdout = oldStdout // Восстанавливаем стандартный вывод
}

func TestHandleKill(t *testing.T) {
	expected := "Missing argument for kill command\n"
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Println(err)
		return
	}
	os.Stdout = w
	handleKill(nil)
	buf := make([]byte, 100)
	n, err := r.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	output := string(buf[:n])
	if output != expected {
		t.Errorf("[%d] returned output - %v - different from expected - %v", 0, output, expected)
	}
	w.Close()
	os.Stdout = oldStdout // Восстанавливаем стандартный вывод
}
