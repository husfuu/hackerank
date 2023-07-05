package main

import "fmt"

func getLocation(tree_loc int32, fruit_dist []int32) []int32 {
	var res []int32
	for _, v := range fruit_dist {
		res = append(res, v+tree_loc)
	}
	return res
}

func getManyFruit(fruit_locs []int32, s int32, t int32) int32 {
	var res int32 = 0
	for _, v := range fruit_locs {
		if v >= s && v <= t {
			res += 1
		}
	}
	return res
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// s: starting point of Sam's house location.
	// t: ending location of Sam's house location.
	// a: location of the Apple tree.
	// b: location of the Orange tree.
	// apples: distances at which each apple falls from the tree.
	// oranges: distances at which each orange falls from the tree.
	apple_locs := getLocation(a, apples)
	fmt.Println("apple loc: ", apple_locs)
	orange_locs := getLocation(b, oranges)
	fmt.Println("orange loc: ", orange_locs)

	apple_fruit := getManyFruit(apple_locs, s, t)
	orange_fruit := getManyFruit(orange_locs, s, t)

	fmt.Println(apple_fruit)
	fmt.Println(orange_fruit)
}

func main() {
	var s int32 = 7
	var t int32 = 11
	var a int32 = 5
	var b int32 = 15
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(s, t, a, b, apples, oranges)
	// fmt.Println(getLocation(res))
}
