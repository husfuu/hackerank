package main

import "testing"

type testStruct struct {
	arg      [][]int32
	expected int32
}

func TestDiagonalDiff(t *testing.T) {
	// test 0
	arr1 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	var expected1 int32 = 2

	var tests = []testStruct{
		{
			arr1,
			expected1,
		},
	}

	for _, test := range tests {
		output := diagonalDifference(test.arg)

		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}

	}

}
