package main

import "testing"

type testStruct struct {
	k      int32
	height []int32
	exp    int32
}

func TestHurdleRace(t *testing.T) {

	tests := []testStruct{
		{
			k:      4,
			height: []int32{1, 6, 3, 5, 2},
			exp:    2,
		},
		{
			k:      7,
			height: []int32{2, 5, 4, 5, 2},
			exp:    0,
		},
	}

	for _, test := range tests {
		output := hurdleRace(test.k, test.height)

		if output != test.exp {
			t.Errorf("Output %d not equal to expected %d", output, test.exp)
		}

	}

}
