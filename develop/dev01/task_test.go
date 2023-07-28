package main

import (
	"testing"
)

type AddressError struct {
	address string
	err     string
}

func TestGetTimeOK(t *testing.T) {
	testCases := []string{"0.beevik-ntp.pool.ntp.org"}

	for index, testCase := range testCases {
		_, err := GetTime(testCase)
		if err != nil {
			t.Errorf("[%d] returned error - %v - different from expected - %v", index, err, nil)
		}
	}
}
func TestGetTimeError(t *testing.T) {
	testCases := []AddressError{
		{
			address: "failAddress",
			err:     "lookup failAddress on 127.0.0.53:53: server misbehaving",
		},
	}
	for index, testCase := range testCases {
		_, err := GetTime(testCase.address)
		if err.Error() != testCase.err {
			t.Errorf("[%d] returned error - %v - different from expected - %v", index, err, testCase)
		}
	}
}
