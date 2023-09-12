package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var min_ch, max_ch int32
	var min, max int32 = scores[0], scores[0]

	for i := 0; i < len(scores); i++ {
		if scores[i] < min {
			fmt.Println(min_ch)
			min = scores[i]
			min_ch += 1
		} else if scores[i] > max {
			max = scores[i]
			max_ch += 1
		}

	}

	return []int32{max_ch, min_ch}
}

func main() {

}
