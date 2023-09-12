package main

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here

	// find maximum
	var max_height int32
	for _, v := range height {
		if v > max_height {
			max_height = v
		}
	}
	if max_height < k {
		return 0
	}
	return max_height - k
}

func main() {

}
