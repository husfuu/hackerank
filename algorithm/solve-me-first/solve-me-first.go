package main

import (
	"fmt"
	"time"
)

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return a + b
}

func main() {

	customDate := "2023-04-26"
	t, err := time.Parse("2006-01-02", customDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customDate)
	fmt.Println(t)

	customDate = "2023-0426"
	t, err = time.Parse("2006-01-02", customDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customDate)
	fmt.Println(t)
}
