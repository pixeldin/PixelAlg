package tree

import "fmt"

/*
  levelOrder 按照层级遍历一颗二叉树,
  由上至下, 每层从左至右
  思路1: 逐层遍历, 使用额外队列(先进先出)把左右子树存储
  思路2: DFS, 深度优先遍历左子树, 每进入一层层数+1, 每层存储一个单独的切片.
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	lv := make([]int, 0)
	lv = append(lv, root.Val)
	ret = append(ret, lv)

	if root.Left != nil {
		zigMarkTreeLvl(root.Left, &ret, 1)
	}
	if root.Right != nil {
		zigMarkTreeLvl(root.Right, &ret, 1)
	}
	return ret
}

func zigMarkTreeLvl(node *TreeNode, ret *[][]int, lvl int) {
	fmt.Println(node.Val)
	if len(*ret) < lvl+1 {
		lv := make([]int, 0)
		*ret = append(*ret, lv)
	}
	if lvl%2 == 0 {
		// 队尾插入
		(*ret)[lvl] = append((*ret)[lvl], node.Val)
	} else {
		// 队头插入
		(*ret)[lvl] = append([]int{node.Val}, (*ret)[lvl]...)
	}
	lvl++

	if node.Left != nil {
		zigMarkTreeLvl(node.Left, ret, lvl)
	}
	if node.Right != nil {
		zigMarkTreeLvl(node.Right, ret, lvl)
	}
}
