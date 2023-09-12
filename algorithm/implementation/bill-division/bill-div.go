package billdivision

import "fmt"

func sum_bill(arr []int32, k int32) int32 {
	var sum int32

	for i := 0; i < len(arr); i++ {
		if int32(i) != k {
			sum += arr[i]
		}
	}

	return sum
}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	sum_bill_actual := sum_bill(bill, k)
	b_actual := sum_bill_actual / 2
	if b_actual == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - b_actual)
	}
}
