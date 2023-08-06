package main

import (
	"bytes"
	"errors"
	"strconv"
	"testing"
)

type ArgsError struct {
	fields    string
	delimiter string
	separated bool
	err       error
}

type FieldError struct {
	fieldIndexStr string
	maxIndex      int
	indexResult   int
	err           error
}

func TestCut(t *testing.T) {
	testCases := []ArgsError{
		{
			fields:    "1",
			separated: false,
			err:       nil,
			delimiter: "\t",
		},
		{
			fields:    "a",
			separated: true,
			err:       strconv.ErrSyntax,
		},
		{
			fields:    "1",
			separated: true,
			err:       nil,
			delimiter: "z",
		},
	}
	for numCases, testCase := range testCases {
		input := "Hello	World"

		// Создаем буфер и инициализируем сканер с помощью буфера
		buffer := bytes.NewBufferString(input)

		err := Cut(buffer, testCase.fields, testCase.delimiter, testCase.separated)
		if !errors.Is(err, testCase.err) {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCases, err, testCase.err)
		}
	}
}

func TestParseFieldIndex(t *testing.T) {
	testCases := []FieldError{
		{
			fieldIndexStr: "10",
			maxIndex:      3,
			err:           nil,
			indexResult:   3,
		},
		{
			fieldIndexStr: "3",
			maxIndex:      10,
			err:           nil,
			indexResult:   2,
		},
	}
	for numCases, testCase := range testCases {
		index, err := parseFieldIndex(testCase.fieldIndexStr, testCase.maxIndex)
		if err != testCase.err {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCases, err, testCase.err)
		}
		if index != testCase.indexResult {
			t.Errorf("[%d] returned index - %v - different from expected - %v", numCases, index, testCase.indexResult)
		}
	}
}

func TestParseFieldIndexes(t *testing.T) {
	testCases := []FieldError{
		{
			fieldIndexStr: "a-b",
			maxIndex:      3,
			err:           strconv.ErrSyntax,
		},
		{
			fieldIndexStr: "1-b",
			maxIndex:      3,
			err:           strconv.ErrSyntax,
		},
		{
			fieldIndexStr: "1-3",
			maxIndex:      10,
			err:           nil,
		},
	}
	for numCases, testCase := range testCases {
		_, err := parseFieldIndexes(testCase.fieldIndexStr, testCase.maxIndex)
		if !errors.Is(err, testCase.err) {
			t.Errorf("[%d] returned err - %v - different from expected - %v", numCases, err, testCase.err)
		}
	}
}
