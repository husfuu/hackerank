package main

import (
	"testing"
)

type testStruct struct {
	arg1     []int32
	arg2     []int32
	expected []int32
}

func TestCompareTripletsTest(t *testing.T) {
	// test 0
	alice0 := []int32{5, 6, 7}
	bob0 := []int32{3, 6, 10}
	res0 := []int32{1, 1}

	// test 1
	alice1 := []int32{17, 28, 30}
	bob1 := []int32{99, 16, 8}
	res1 := []int32{2, 1}

	// test
	alice2 := []int32{1, 2, 3}
	bob2 := []int32{1, 2, 3}
	res2 := []int32{0, 0}

	// test 2
	var tests = []testStruct{
		{
			alice0,
			bob0,
			res0,
		},
		{
			alice1,
			bob1,
			res1,
		},
		{
			alice2,
			bob2,
			res2,
		},
	}

	for _, test := range tests {
		output := compareTriplets(test.arg1, test.arg2)

		for i := 0; i < len(output); i++ {
			if output[i] != test.expected[i] {
				t.Errorf("Output %q not equal to expected %q", output, test.expected)
			}
		}
	}
}
