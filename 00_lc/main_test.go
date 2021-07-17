package main

import (
	arr "PixelAlg/00_LeetCode/array"
	string2 "PixelAlg/00_LeetCode/string"
	"PixelAlg/00_LeetCode/tree"
	"PixelAlg/00_LeetCode/two_pointer"
	"fmt"
	"testing"
)

func TestSome(t *testing.T) {
	//test := []int{0, -1, 1, -4, 2, 3}
	//sum := ThreeSumV(test)
	//fmt.Println(sum)

	test2 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	sum2 := arr.ThreeSumV(test2)
	fmt.Println(sum2)

}

func TestThreeSum(t *testing.T) {
	//test2 := []int{0, 2, -1, 1, 0, 1, -1, -3, 3, -2, 0}
	//test1 := []int{-1,0,1,2,-1,-4}
	test1 := []int{0, 0, 0, 0}

	sum2 := arr.ThreeSumLocal(test1)
	fmt.Print(sum2)
}

func TestThreeSumClosestlocal(t *testing.T) {
	arr.ThreeSumClosestLocal()
}

func TestString(t *testing.T) {
	string2.TestLocal()
}

func TestTree(t *testing.T) {
	//tree.TestMinTravel()
	//tree.TestValidBST()
	//tree.TestisSameTree()
	tree.TestMaxDepth()
}

func TestDuplicates(t *testing.T) {
	// 双指针去重
	//two_pointer.RemoveDuplicates([]int{0,0,1,1,1,2,2,3,3,4})

	// 双指针去除指定
	//two_pointer.RemoveElement([]int{3, 2, 2, 3}, 3)
	// 移除0元素, 保持相对顺序(稳定)
	two_pointer.MoveZeroes([]int{0, 0, 1, 1, 0, 2, 0, 3, 3, 0})
}
