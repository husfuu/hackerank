package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	greater := 0
	greater_freq := 0

	for i := 0; i < len(candles); i++ {
		if candles[i] > int32(greater) {
			greater = int(candles[i])
		}
	}

	for i := 0; i < len(candles); i++ {
		if candles[i] == int32(greater) {
			greater_freq += 1
		}
	}

	return int32(greater_freq)
}

func main() {
	candles := []int32{4, 4, 1, 3}
	candles2 := []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(candles))
	fmt.Println(birthdayCakeCandles(candles2))
}
