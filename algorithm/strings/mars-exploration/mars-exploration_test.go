package main

import "testing"

type testStruct struct {
	s   string
	exp int32
}

func TestMarsExploration(t *testing.T) {
	tests := []testStruct{
		{
			s:   "SOSSOT",
			exp: 1,
		},
		{
			s:   "SOSSOSSOS",
			exp: 0,
		},
		{
			s:   "SOSOSOSOSDSDSKWOSDOSDOASDOASDFAFJDFDOSOSOWNSOSOSNDSKDDOSOSOSJDSDSOOSOSDSDOSASSOASDSAOSOSODSDSOASDWS",
			exp: 67,
		},
	}

	for _, v := range tests {

		output := marsExploration(v.s)
		if output != v.exp {
			t.Errorf("Output %d not equal to expected %d", output, v.exp)
		}

	}
}
