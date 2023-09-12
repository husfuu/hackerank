package main

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var res int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			sum := ar[i] + ar[j]
			if sum%k == 0 {
				res += 1
			}
		}
	}

	return res
}

func main() {

}
