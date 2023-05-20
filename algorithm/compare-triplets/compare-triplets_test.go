package main

type testStruct struct {
	arg1     []int32
	arg2     []int32
	expected []int32
}

func compareTripletsTest(t *testing.T) {
	alice1 := []int32{17, 28, 30}
	bob1 := []int32{99, 16, 8}
	res1 := []int32{1, 2}
	var tests = []testStruct{
		testStruct{
			alice1,
			bob1,
			res1,
		},
	}

	for _, test := range tests {
		// output := compareTriplets(test.arg1, test.arg2)

	}
}
