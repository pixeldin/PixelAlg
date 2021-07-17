package dp

import (
	"PixelAlg/00_LeetCode/tree"
	"fmt"
	"log"
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

func TestRob(t *testing.T) {
	i := rob([]int{4, 1, 7, 3, 8})
	fmt.Println(i)

	//i = rob2([]int{4, 1, 7, 3, 8})
	i = rob2([]int{200, 3, 140, 20, 10})
	fmt.Println(i)
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

func buildLocalTree() *tree.TreeNode {
	a := tree.BuildTreeNode(7)
	b := tree.BuildTreeNode(30)
	c := tree.BuildTreeNode(5)
	d := tree.BuildTreeNode(1)
	e := tree.BuildTreeNode(3)

	f := tree.BuildTreeNode(200)
	g := tree.BuildTreeNode(2)
	h := tree.BuildTreeNode(4)
	i := tree.BuildTreeNode(6)
	j := tree.BuildTreeNode(9)
	k := tree.BuildTreeNode(8)
	l := tree.BuildTreeNode(10)

	a.Left = b
	a.Right = c

	b.Left = d
	b.Right = e

	//
	c.Left = nil
	c.Right = f

	d.Left = g
	d.Right = h

	e.Left = i
	e.Right = j

	f.Left = k
	f.Right = l

	return a
}

func TestRobTree(t *testing.T) {
	//rob3TraceAll(buildLocalTree(), 0)
	record := make(map[*tree.TreeNode]int, 0)
	ret := rob3WithRecord(buildLocalTree(), record)
	log.Print(ret)
}
