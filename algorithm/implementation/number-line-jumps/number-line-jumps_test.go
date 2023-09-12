package main

import (
	"fmt"
	"testing"
)

type testStruct struct {
	x1  int32
	v1  int32
	x2  int32
	v2  int32
	exp string
}

func TestKangaroo(t *testing.T) {
	var tests = []testStruct{
		{
			x1:  2,
			v1:  1,
			x2:  1,
			v2:  2,
			exp: "YES",
		},
		{
			x1:  0,
			v1:  3,
			x2:  4,
			v2:  2,
			exp: "YES",
		},
		{
			x1:  0,
			v1:  2,
			x2:  5,
			v2:  3,
			exp: "NO",
		},
		{
			x1:  21,
			v1:  6,
			x2:  47,
			v2:  3,
			exp: "NO",
		},
		{
			x1:  4523,
			v1:  8092,
			x2:  9419,
			v2:  8076,
			exp: "YES",
		},
		{
			x1:  43,
			v1:  2,
			x2:  70,
			v2:  2,
			exp: "NO",
		},
	}

	for _, test := range tests {
		output := kangaroo(test.x1, test.v1, test.x2, test.v2)

		if output != test.exp {
			fmt.Println(test.x1)
			t.Errorf("Output %s not equal to expected %s", output, test.exp)
		}
	}
}
