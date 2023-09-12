package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here

	if x1 < x2 && v1 <= v2 {
		return "NO"
	}

	// if v1 == v2 {
	// 	return "NO"
	// }

	j := (x2 - x1) % (v1 - v2)

	if j == 0 {
		return "YES"
	}
	return "NO"

	// fmt.Println(j)

	// // for x1 <= 10000 && x2 <= 10000 {
	// // 	x1 = x1 + v1
	// // 	x2 = x2 + v2
	// // 	// fmt.Println("x1: ", x1)
	// // 	// fmt.Println("x2: ", x2)
	// // 	if x1 == x2 {
	// // 		return "YES"
	// // 	}
	// // }

	// for {
	// 	x1 = x1 + v1
	// 	x2 = x2 + v2
	// 	// fmt.Println("x1: ", x1)
	// 	// fmt.Println("x2: ", x2)
	// 	if x1 == x2 {
	// 		fmt.Println("x1: ", x1)
	// 		fmt.Println("x2: ", x2)
	// 		return "YES"
	// 	}
	// }

	// return "NO"
}

func main() {
	res := kangaroo(4523, 8092, 9419, 8076)
	// res1 := kangaroo(0, 3, 4, 2)
	// res2 := kangaroo(0, 2, 5, 3)

	fmt.Println(res)

	// fmt.Println(res1)
	// fmt.Println(res2)
}
