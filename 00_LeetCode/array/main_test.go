package main

import (
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	test := []int{0, -1, 1, -4, 2, 3}
	sum := ThreeSumV(test)
	fmt.Print(sum)

	test2 := []int{0, 0, -1, 1, 0, 1, -1, -3, 3}
	sum2 := ThreeSumV(test2)
	fmt.Print(sum2)

}
