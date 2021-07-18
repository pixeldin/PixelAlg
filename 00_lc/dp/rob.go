package dp

import (
	"PixelAlg/00_lc/tree"
	"PixelAlg/00_lc/util"
	"fmt"
)

// 198 相邻房间选择最大收益
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] > nums[1] {
		nums[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		nums[i] = util.Max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

// 213. House Robber II
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//opt1, opt2 := 0, 0
	o1Num := nums[:len(nums)-1]
	o2Num := nums[1:]
	co2 := make([]int, len(o2Num)) //使用copy函数必须复制切片的结构必须和源数据结构一致
	copy(co2, o2Num)
	opt1 := rob(o1Num)
	opt2 := rob(co2)
	return util.Max(opt1, opt2)
}

func printTab(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("--")
	}
	fmt.Print("|")
}

// 337. House Robber III
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var PrintLvl = true

// 递归,穷举
func rob3TraceAll(root *tree.TreeNode, lvl int) int {
	if root == nil {
		return 0
	}
	var ret = 0
	if PrintLvl {
		lvl++
		printTab(lvl)
		fmt.Println(" ", root.Val)
	}
	if root.Left != nil {
		// 此处的 root.Left.Left会与下一层的root.left重复计算,所以可以考虑填入备忘录
		ret += rob3TraceAll(root.Left.Left, lvl)
		ret += rob3TraceAll(root.Left.Right, lvl)
	}
	if root.Right != nil {
		ret += rob3TraceAll(root.Right.Left, lvl)
		ret += rob3TraceAll(root.Right.Right, lvl)
	}
	// 重复处
	return util.Max(root.Val+ret, rob3TraceAll(root.Left, lvl)+rob3TraceAll(root.Right, lvl))
}

// 带有备忘录的递归
func rob3WithRecord(root *tree.TreeNode, record map[*tree.TreeNode]int) int {
	if root == nil {
		return 0
	}
	var ret = 0
	if _, ok := record[root]; ok {
		return record[root]
	}
	//if PrintLvl {
	//	lvl++
	//	printTab(lvl)
	//	fmt.Println(" ", root.Val)
	//}
	if root.Left != nil {
		// 此处的 root.Left.Left会与下一层的root.left重复计算,所以可以考虑填入备忘录
		ret += rob3WithRecord(root.Left.Left, record)
		ret += rob3WithRecord(root.Left.Right, record)
	}
	if root.Right != nil {
		ret += rob3WithRecord(root.Right.Left, record)
		ret += rob3WithRecord(root.Right.Right, record)
	}

	// mark record
	ret = util.Max(root.Val+ret, rob3WithRecord(root.Left, record)+rob3WithRecord(root.Right, record))
	record[root] = ret
	return ret
}
