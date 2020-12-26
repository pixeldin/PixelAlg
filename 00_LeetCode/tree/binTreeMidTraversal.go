package tree

import (
	"fmt"
)

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestMinTravel() {
	tr := buildTree()
	traversal := inorderTraversal(tr)
	fmt.Println(traversal)
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	midTraversal(root, &ret)
	return ret
}

func midTraversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		midTraversal(root.Left, ret)
	}
	//midTraversal(root, ret)
	*ret = append(*ret, root.Val)
	if root.Right != nil {
		midTraversal(root.Right, ret)
	}
}

func buildTree() *TreeNode {
	left := TreeNode{
		3,
		nil,
		nil,
	}
	right := TreeNode{
		2,
		&left,
		nil,
	}
	root := TreeNode{
		2,
		nil,
		&right,
	}
	return &root
}
