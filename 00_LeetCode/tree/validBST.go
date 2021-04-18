package tree

import (
	"fmt"
	"math"
)

func TestValidBST() {
	tr := BuildTreeWithThree(2, 1, 3)
	fmt.Println(isValidBST(tr))
}

func isValidBST(root *TreeNode) bool {
	min := math.Inf(-1)
	max := math.Inf(1)
	return legalCheck(root, min, max)
}

// 递归子树判定
func legalCheck(tar *TreeNode, min, max float64) bool {
	if tar == nil {
		return true
	}
	val := float64(tar.Val)
	if val < min {
		return false
	}
	if val > max {
		return false
	}
	return legalCheck(tar.Left, min, val) && legalCheck(tar.Right, val, max)
}

func isValidBST_(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidBST_(root.Left, min, v) && isValidBST_(root.Right, v, max)
}
