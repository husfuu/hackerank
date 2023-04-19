package main

import "fmt"

func bubleSort(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func minMaxSum(arr []int32) {
	var sorted = bubleSort(arr)
	var resMin = 0
	var resMax = 0

	for i := 0; i < len(sorted); i++ {
		if i < (len(sorted) - 1) {
			resMin += int(sorted[i])
		}
		if i > 0 {
			resMax += int(sorted[i])
		}
	}

	fmt.Print(resMin, resMax)
}

func main() {
	var numbs = []int32{7, 69, 2, 221, 8974}

	// var res = bubleSort(numbs)

	// fmt.Println(res)
	minMaxSum(numbs)
	// fmt.Println(res)
}
