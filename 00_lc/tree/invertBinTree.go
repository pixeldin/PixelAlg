package tree

// 反转一颗二叉树
func invertTree(root *TreeNode) *TreeNode {
	//if root.Left == nil || root.Right == nil {
	//	return nil
	//}
	if root == nil {
		return root
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
