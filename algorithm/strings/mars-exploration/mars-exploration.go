package main

import (
	"fmt"
)

func marsExploration(s string) int32 {
	// Write your code here
	var res int32
	for i := 0; i < len(s); i += 3 {

		if fmt.Sprintf("%c", s[i]) != "S" {
			res += 1
		}
		if fmt.Sprintf("%c", s[i+1]) != "O" {
			res += 1
		}
		if fmt.Sprintf("%c", s[i+2]) != "S" {
			res += 1
		}

	}
	return res
}

func main() {
	for _, v := range "Satuiiii" {
		fmt.Println(v == 'S')
	}
	word := "Satuiiii"
	for i := 0; i < len(word); i++ {
		r := word[i]

		r1 := fmt.Sprintf("%c", r)
		fmt.Println(r1)

		fmt.Println((r1 == "S"))
	}

}
