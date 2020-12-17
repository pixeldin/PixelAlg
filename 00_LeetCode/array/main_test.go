package main

import (
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	test := []int{0, -1, 1, -4, 2, 3}
	sum := ThreeSum(test)
	fmt.Print(sum)
}
