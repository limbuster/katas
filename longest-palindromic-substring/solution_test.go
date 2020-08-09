package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input, output string
	}{
		{input: "", output: ""},
		{input: "a", output: "a"},
		{input: "bb", output: "bb"},
		{input: "babad", output: "bab"},
		{input: "cbbd", output: "bb"},
	}

	for _, testCase := range testCases {
		got := LongestPalindrome(testCase.input)
		want := testCase.output

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{input: "", output: true},
		{input: "a", output: true},
		{input: "babad", output: false},
		{input: "cbbd", output: false},
		{input: "aba", output: true},
		{input: "abcccba", output: true},
		{input: "abccfba", output: false},
	}

	for _, testCase := range testCases {
		got := IsPalindrome(testCase.input)
		want := testCase.output

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}
}
