package main

import "testing"

type testStruct struct {
	scores []int32
	exp    [2]int32
}

func TestBreakingRecords(t *testing.T) {
	var tests = []testStruct{
		{
			scores: []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42},
			exp:    [2]int32{4, 0},
		},
		{
			scores: []int32{10, 5, 20, 20, 4, 5, 2, 25, 1},
			exp:    [2]int32{2, 4},
		},
	}

	for _, test := range tests {
		output := breakingRecords(test.scores)

		if test.exp != [2]int32(output) {
			t.Errorf("Output %v not equal to expected %v", output, test.exp)
		}
	}
}
