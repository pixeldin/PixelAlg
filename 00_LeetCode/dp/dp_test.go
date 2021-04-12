package dp

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	stairs := ClimbStairs(35)
	fmt.Print(stairs)
}

func TestIsSubsequence1(t *testing.T) {
	//fmt.Println(IsSubsequence1("abcd", "sabc123abcd"))
	fmt.Println(getFibSum(2))
	fmt.Println(getFibSum(5))
	fmt.Println(GotFibSumWithMark(2))
	fmt.Println(GotFibSumWithMark(5))
}

func TestPath(t *testing.T) {
	paths := UniquePaths(3, 7)
	fmt.Println(paths)
}

func TestSome(t *testing.T) {
	maxSubArray_l([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func maxSubArray_l(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		max = Max(max, sum)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
