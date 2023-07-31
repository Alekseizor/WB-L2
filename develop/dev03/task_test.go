package main

import (
	"testing"
)

type RowError struct {
	row    string
	unique bool
	err    error
}

type RowErrorStr struct {
	row    string
	unique bool
	err    string
}

type RowsError struct {
	row     [][]string
	reverse bool
	err     error
}

type RowsRows struct {
	row       [][]string
	rowResult [][]string
	n         bool
	k         int
}

func TestReadLinesOK(t *testing.T) {
	testCases := []RowError{
		{
			row:    "text",
			unique: true,
			err:    nil,
		},
	}
	for numCase, testCase := range testCases {
		_, err := readLines(testCase.row, testCase.unique)
		if err != testCase.err {
			t.Errorf("[%d] returned error - %v - different from expected - %v", numCase, err, testCase.err)
		}
	}
}

func TestReadLinesError(t *testing.T) {
	testCases := []RowErrorStr{
		{
			row:    "textFalse",
			unique: false,
			err:    "open textFalse: no such file or directory",
		},
	}
	for numCase, testCase := range testCases {
		_, err := readLines(testCase.row, testCase.unique)
		if err.Error() != testCase.err {
			t.Errorf("[%d] returned error - %v - different from expected - %v", numCase, err, testCase.err)
		}
	}
}

func TestWriteLinesOK(t *testing.T) {
	testCases := []RowsError{
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			reverse: true,
			err:     nil,
		},
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			reverse: false,
			err:     nil,
		},
	}
	for numCase, testCase := range testCases {
		err := writeLines(testCase.row, testCase.reverse)
		if err != testCase.err {
			t.Errorf("[%d] returned error - %v - different from expected - %v", numCase, err, testCase.err)
		}
	}
}

func TestSortLineOK(t *testing.T) {
	testCases := []RowsRows{
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			n: false,
			rowResult: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			k: 0,
		},
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			n: false,
			rowResult: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			k: 13,
		},
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			n: true,
			rowResult: [][]string{
				{
					"b",
					"1",
				},
				{
					"a",
					"2",
				},
			},
			k: 1,
		},
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			n: true,
			rowResult: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			k: 0,
		},
		{
			row: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			n: true,
			rowResult: [][]string{
				{
					"a",
					"2",
				},
				{
					"b",
					"1",
				},
			},
			k: 13,
		},
		{
			row: [][]string{
				{
					"a",
					"1",
				},
				{
					"b",
					"2",
				},
			},
			n: true,
			rowResult: [][]string{
				{
					"a",
					"1",
				},
				{
					"b",
					"2",
				},
			},
			k: 1,
		},
	}
	for numCase, testCase := range testCases {
		result := sortLine(testCase.row, testCase.k, testCase.n)
		for numRow, row := range result {
			for numWord, word := range row {
				if word != testCase.rowResult[numRow][numWord] {
					t.Errorf("[%d] returned word - %v - different from expected - %v", numCase, word, testCase.rowResult[numRow][numWord])
				}
			}
		}
	}
}
