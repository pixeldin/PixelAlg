package tree

/*
  levelOrder 按照层级遍历一颗二叉树,
  由上至下, 每层从左至右
  思路1: 逐层遍历, 使用额外队列(先进先出)把左右子树存储
  思路2: DFS, 深度优先遍历左子树, 每进入一层层数+1, 每层存储一个单独的切片.
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	lv := make([]int, 0)
	lv = append(lv, root.Val)
	ret = append(ret, lv)
	if root.Left != nil {
		markTreeLvl(root.Left, &ret, 1)
	}
	if root.Right != nil {
		markTreeLvl(root.Right, &ret, 1)
	}
	return ret
}

func markTreeLvl(node *TreeNode, ret *[][]int, lvl int) {
	//fmt.Println(node.Val)
	if len(*ret) < lvl+1 {
		lv := make([]int, 0)
		*ret = append(*ret, lv)
	}
	(*ret)[lvl] = append((*ret)[lvl], node.Val)
	lvl++
	if node.Left != nil {
		markTreeLvl(node.Left, ret, lvl)
	}
	if node.Right != nil {
		markTreeLvl(node.Right, ret, lvl)
	}
}
