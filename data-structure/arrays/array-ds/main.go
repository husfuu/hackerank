package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// Write your code here
	var res []int32
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}

	return res
}

func main() {
	arr := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(arr))
}
