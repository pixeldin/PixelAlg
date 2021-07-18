package tree

import (
	"PixelAlg/00_lc/util"
)

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return util.Max(getHeight(root.Right), getHeight(root.Left)) + 1
}

// 求最大高度-最小高度
// |左子树 - 右子树| 小于 1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := getHeight(root.Left)
	rh := getHeight(root.Right)
	absMine := lh - rh
	if absMine < 0 {
		absMine *= -1
	}
	return absMine < 2 && isBalanced(root.Right) && isBalanced(root.Left)
}

func deepCheck(current int, min, max *int, node *TreeNode) bool {
	if node == nil {
		return true
	}

	current++
	if current > *max {
		*max = current
	}
	if *max-*min > 1 {
		return false
	}
	// 根节点一个分支为空 或者 非根节点到底部了
	if (current == 0 && (node.Left == nil || node.Right == nil)) ||
		(node.Left == nil && node.Right == nil) {
		if current < *min {
			*min = current
		}
		if *max-*min > 1 {
			return false
		}
		return true
	}
	return deepCheck(current, min, max, node.Right) &&
		deepCheck(current, min, max, node.Left)
}
