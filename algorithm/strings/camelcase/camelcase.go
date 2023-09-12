package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	var numWord int32 = 1
	for _, char := range s {
		if unicode.IsUpper(char) {
			numWord += 1
		}
	}

	return numWord
}

func main() {
	res := camelcase("hebatBanget")
	fmt.Println(res)
}
