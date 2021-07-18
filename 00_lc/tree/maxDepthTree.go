package tree

import (
	"PixelAlg/00_lc/util"
	"fmt"
)

func TestMaxDepth() {
	three := BuildTreeWithThree(1, 3, 3)
	fmt.Println(maxDepth(three))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return util.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
