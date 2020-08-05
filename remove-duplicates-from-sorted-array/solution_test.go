package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	isSliceEqual := func(a []int, b []int) bool {

		for i, v := range a {
			if len(b) == 0 {
				return false
			}
			if v != b[i] {
				return false
			}
		}
		return true
	}

	cases := []struct {
		input   []int
		output  int
		output2 []int
	}{
		{input: []int{}, output: 0, output2: []int{}},
		{input: []int{1, 1, 1}, output: 1, output2: []int{1}},
		{input: []int{1, 1, 2}, output: 2, output2: []int{1, 2}},
		{input: []int{0, 1, 2, 3, 4}, output: 5, output2: []int{0, 1, 2, 3, 4}},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, output: 5, output2: []int{0, 1, 2, 3, 4}},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4}, output: 5, output2: []int{0, 1, 2, 3, 4}},
	}

	for _, testCase := range cases {
		got := RemoveDuplicates(testCase.input)
		want := testCase.output

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}

		if !isSliceEqual(testCase.output2, testCase.input[:got]) {
			t.Errorf("want %+v got %+v", testCase.output2, testCase.input[:got])
		}
	}
}
