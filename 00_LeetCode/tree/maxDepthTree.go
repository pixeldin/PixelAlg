package tree

import "fmt"

func TestMaxDepth()  {
	three := BuildTreeWithThree(1, 3, 3)
	fmt.Println(maxDepth(three))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}