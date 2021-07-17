package tree

// isSymmetric 思路1: 递归判定左右子树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	//return check(root.Right, root.Left)
	// 从根节点开始, 防止根节点为nil, 访问其左右节点报空指针
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if q != nil && p != nil {
		if q.Val != p.Val {
			return false
		}
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// isSymmetricV2 思路2: 反转左子树并与右子树判断是否相同
func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(invertTree_(root.Left), root.Right)
}

func invertTree_(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree_(root.Left)
	invertTree_(root.Right)
	return root
}
