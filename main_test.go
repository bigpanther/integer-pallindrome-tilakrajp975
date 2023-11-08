package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input          int
		expectedOutput bool
	}{
		{1, true},
		{-1, true},
		{11, true},
		{12, false},
		{112, false},
		{121, true},
		{-50, false},
		{-8297, false},
		{0, true},
		{1211111111, false},
		{1211111121, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			actual := IsPalindrome(test.input)
			if test.expectedOutput != actual {
				t.Fail()
			}
		})
	}
}
