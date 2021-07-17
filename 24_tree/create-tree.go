package _4_tree

import (
	binarytree "PixelAlg/23_binarytree"
)

// 前序遍历, -1为空节点
func CreateBinTree(lt *[]int) *binarytree.TreeNode {
	if len(*lt) == 0 {
		return nil
	}
	var node *binarytree.TreeNode
	v := (*lt)[0]
	*lt = (*lt)[1:]
	if v == -1 {
		// reach bottom
		return node
	}
	node = &binarytree.TreeNode{Val: v}
	node.Left = CreateBinTree(lt)
	node.Right = CreateBinTree(lt)
	return node
}
