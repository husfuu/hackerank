package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var NUM_LOOP = 3
	result := []int32{0, 0}
	for i := 0; i < NUM_LOOP; i++ {
		if a[i] > b[i] {
			result[0] += 1
		} else if a[i] < b[i] {
			result[1] += 1
		}
	}
	return result
}

func main() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}

	result := compareTriplets(a, b)

	fmt.Println(result)
}
