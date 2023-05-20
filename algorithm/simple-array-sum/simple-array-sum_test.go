package main

import "testing"

type testStruct struct {
	arg      []int32
	expected int32
}

func TestSimpleArraySum(t *testing.T) {
	// test 0
	arr1 := []int32{1, 2, 3, 4, 10, 11}
	var expected1 int32 = 31

	var tests = []testStruct{
		{
			arr1,
			expected1,
		},
	}

	for _, test := range tests {
		output := simpleArraySum(arr1)

		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
