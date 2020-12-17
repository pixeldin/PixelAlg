package main

import "fmt"

func main() {
	test := []int{1, 9, 3, 4, 8, 10}
	s1 := TwoSum(test, 18)
	fmt.Print(s1)
	s2 := BetterTwoSum(test, 10)
	fmt.Print(s2)
	s2 = BetterTwoSum(test, 18)
	fmt.Print(s2)
}
