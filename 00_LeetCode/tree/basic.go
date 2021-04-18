package tree

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeWithThree(rootVal, leftVal, rightVal int) *TreeNode {
	left := TreeNode{
		leftVal,
		nil,
		nil,
	}
	right := TreeNode{
		rightVal,
		nil,
		nil,
	}
	root := TreeNode{
		rootVal,
		&left,
		&right,
	}
	return &root
}

func BuildTreeNode(rootVal int) *TreeNode {
	root := TreeNode{
		rootVal,
		nil,
		nil,
	}
	return &root
}
