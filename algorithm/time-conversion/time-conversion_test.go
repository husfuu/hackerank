package main

import "testing"

type testStruct struct {
	arg1     string
	expected string
}

func TestTimeConversionTest(t *testing.T) {
	// test 0
	input1 := "07:05:45PM"
	exp1 := "19:05:45"

	input2 := "12:40:22AM"
	exp2 := "00:40:22"

	input3 := "06:40:03AM"
	exp3 := "06:40:03"

	input4 := "12:00:00AM"
	exp4 := "00:00:00"

	input5 := "11:59:59PM"
	exp5 := "23:59:59"

	var tests = []testStruct{
		{
			input1,
			exp1,
		},
		{
			input2,
			exp2,
		},
		{
			input3,
			exp3,
		},
		{
			input4,
			exp4,
		},
		{
			input5,
			exp5,
		},
	}

	for _, test := range tests {
		output := timeConversions(test.arg1)

		for i := 0; i < len(output); i++ {
			if output[i] != test.expected[i] {
				t.Errorf("Output %q not equal to expected %q", output, test.expected)
			}
		}
	}
}
