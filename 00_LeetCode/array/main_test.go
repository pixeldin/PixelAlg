package main

import (
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	//test := []int{0, -1, 1, -4, 2, 3}
	//sum := ThreeSumV(test)
	//fmt.Println(sum)

	test2 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	sum2 := ThreeSumV(test2)
	fmt.Println(sum2)

}

func TestThreeSum(t *testing.T) {
	//test2 := []int{0, 2, -1, 1, 0, 1, -1, -3, 3, -2, 0}
	//test1 := []int{-1,0,1,2,-1,-4}
	test1 := []int{0, 0, 0, 0}

	sum2 := ThreeSumLocal(test1)
	fmt.Print(sum2)
}

func TestThreeSumClosestlocal(t *testing.T) {
	ThreeSumClosestLocal()
}
