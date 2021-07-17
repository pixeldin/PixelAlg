package main

import (
	arr "PixelAlg/00_LeetCode/array"
	"fmt"
)

func main() {
	test := []int{1, 9, 3, 4, 8, 10}
	s1 := arr.TwoSum(test, 18)
	fmt.Print(s1)
	s2 := arr.BetterTwoSum(test, 10)
	fmt.Print(s2)
	s2 = arr.BetterTwoSum(test, 18)
	fmt.Print(s2)
}
