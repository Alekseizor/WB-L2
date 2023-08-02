package main

import "testing"

type WordsAnagram struct {
	words   *[]string
	anagram *map[string][]string
}

func TestSortStringOK(t *testing.T) {
	result := sortString("lexa")
	if result != "aelx" {
		t.Errorf("[%d] returned str - %v - different from expected - %v", 0, result, "aelx")
	}
}

func TestFindAnagramSetsOK(t *testing.T) {
	testCases := []WordsAnagram{
		{
			words: &[]string{"пятка", "пятак", "тяпка", "тяпка", "слиток", "слиток"},
			anagram: &map[string][]string{
				"пятка": {"пятак", "пятка", "тяпка"},
			},
		},
	}
	for numCase, testCase := range testCases {
		result := FindAnagramSets(testCase.words)
		for key, value := range *result {
			words, exists := (*testCase.anagram)[key]
			if !exists {
				t.Errorf("[%d] returned key - %v - different from expected", numCase, key)
			}
			for numWord, word := range value {
				if word != words[numWord] {
					t.Errorf("[%d] returned word - %v - different from expected", numCase, word)
				}
			}
		}
	}
}
