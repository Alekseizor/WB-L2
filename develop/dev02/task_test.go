package dev02

import "testing"

type RowError struct {
	row string
	err error
}

func TestUnpackString(t *testing.T) {
	testCases := []RowError{
		{
			row: "a4bc2\\42d5e",
			err: nil,
		},
		{
			row: "45",
			err: invalidStringError,
		},
	}
	for numCase, testCase := range testCases {
		_, err := UnpackString(testCase.row)
		if err != testCase.err {
			t.Errorf("[%d] returned error - %v - different from expected - %v", numCase, err, testCase.err)
		}
	}
}
