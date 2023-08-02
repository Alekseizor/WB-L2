package main

import (
	"log"
	"os"
	"testing"
)

type ProcessData struct {
	lines      []*RowFlag
	pattern    string
	after      int
	before     int
	context    int
	ignoreCase bool
	fixed      bool
	count      bool
	invert     bool
	lineNum    bool
	counter    int
	err        error
}

func TestReadFileOK(t *testing.T) {
	_, err := readFile("text")
	if err != nil {
		t.Errorf("[%d] returned err - %v - different from expected - %v", 0, err, nil)
	}
}

func TestReadFileError(t *testing.T) {
	_, err := readFile("fail")
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("[%d] returned err - %v - different from expected - %v", 0, err, "open fail: no such file or directory")
	}
}

func TestPrintLinesOK(t *testing.T) {
	lines := []*RowFlag{
		{
			row:    "test",
			output: false,
		},
		{
			row:    "test2",
			output: true,
		},
	}
	expected := "1: test\n"
	// Запоминаем текущий стандартный вывод
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Println(err)
		return
	}
	os.Stdout = w

	// Вызываем функцию, которую нужно протестировать
	printLines(lines, true)
	printLines(lines, false)

	// Читаем вывод из канала
	w.Close()
	out := make([]byte, 1024)
	n, _ := r.Read(out)
	os.Stdout = old
	// Проверяем ожидаемый вывод
	if string(out[:n]) != expected {
		t.Errorf("[%d] the information output is - %v - different from the expected - %v", 0, string(out[:n]), expected)
	}
}

func TestProcessFileOK(t *testing.T) {
	testCases := []ProcessData{
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern: "test",
			invert:  true,
			counter: 0,
		},
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			fixed:   true,
			pattern: "test",
			counter: 2,
			count:   true,
		},
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern:    "test",
			ignoreCase: true,
		},
	}
	for numCase, testCase := range testCases {
		counter, err := processFile(testCase.lines, testCase.pattern, testCase.after, testCase.before, testCase.context, testCase.ignoreCase, testCase.fixed, testCase.count, testCase.invert, testCase.lineNum)
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCase, err, testCase.err)
		}
		if counter != testCase.counter {
			t.Errorf("[%d] returned number - %v - different from expected - %v", numCase, counter, testCase.counter)
		}
	}
}

func TestProcessFileContextOK(t *testing.T) {
	testCases := []ProcessData{
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern: "test",
			counter: 0,
			context: 1,
		},
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern: "test",
			counter: 0,
			context: 1234,
		},
		{
			lines: []*RowFlag{
				{
					row:    "row",
					output: false,
				},
				{
					row:    "row",
					output: false,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern: "test",
			counter: 0,
			context: 1,
		},
		{
			lines: []*RowFlag{
				{
					row:    "test2",
					output: true,
				},
				{
					row:    "row",
					output: false,
				},
				{
					row:    "row",
					output: false,
				},
			},
			pattern: "test",
			counter: 0,
			context: 1,
		},
	}
	for numCase, testCase := range testCases {
		counter, err := processFile(testCase.lines, testCase.pattern, testCase.after, testCase.before, testCase.context, testCase.ignoreCase, testCase.fixed, testCase.count, testCase.invert, testCase.lineNum)
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCase, err, testCase.err)
		}
		if counter != testCase.counter {
			t.Errorf("[%d] returned number - %v - different from expected - %v", numCase, counter, testCase.counter)
		}
	}
}

func TestProcessFileBeforeOK(t *testing.T) {
	testCases := []ProcessData{
		{
			lines: []*RowFlag{
				{
					row:    "test2",
					output: false,
				},
				{
					row:    "row",
					output: true,
				},
				{
					row:    "test2",
					output: true,
				},
			},
			pattern: "test",
			counter: 0,
			before:  1,
		},
	}
	for numCase, testCase := range testCases {
		counter, err := processFile(testCase.lines, testCase.pattern, testCase.after, testCase.before, testCase.context, testCase.ignoreCase, testCase.fixed, testCase.count, testCase.invert, testCase.lineNum)
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCase, err, testCase.err)
		}
		if counter != testCase.counter {
			t.Errorf("[%d] returned number - %v - different from expected - %v", numCase, counter, testCase.counter)
		}
	}
}

func TestProcessFileAfterOK(t *testing.T) {
	testCases := []ProcessData{
		{
			lines: []*RowFlag{
				{
					row:    "test",
					output: false,
				},
				{
					row:    "row",
					output: true,
				},
				{
					row:    "test",
					output: true,
				},
			},
			pattern: "test",
			counter: 0,
			after:   1,
		},
	}
	for numCase, testCase := range testCases {
		counter, err := processFile(testCase.lines, testCase.pattern, testCase.after, testCase.before, testCase.context, testCase.ignoreCase, testCase.fixed, testCase.count, testCase.invert, testCase.lineNum)
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCase, err, testCase.err)
		}
		if counter != testCase.counter {
			t.Errorf("[%d] returned number - %v - different from expected - %v", numCase, counter, testCase.counter)
		}
	}
}
