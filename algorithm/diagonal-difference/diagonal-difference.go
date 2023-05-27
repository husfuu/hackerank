package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	var sum_right int32 = 0
	var sum_left int32 = 0
	var j_min int32 = int32(len(arr))
	for i := 0; i < len(arr); i++ {
		j_min -= 1
		for j := 0; j < len(arr); j++ {
			if i == j {
				sum_right += arr[i][j]
			}
			if j == int(j_min) {
				sum_left += arr[i][j]
			}
		}
	}
	diff := sum_right - sum_left
	diff_abs := math.Abs(float64(diff))
	return int32(diff_abs)
}

func main() {
	a := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	result := diagonalDifference(a)

	fmt.Println(result)
}
