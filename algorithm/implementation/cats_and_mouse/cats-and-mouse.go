package main

import "math"

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	if math.Abs(float64(x-z)) < math.Abs(float64(y-z)) {
		return "Cat A"
	} else if math.Abs(float64(x-z)) == math.Abs(float64(y-z)) {
		return "Mouse C"
	} else {
		return "Cat B"
	}
}

func main() {

}
