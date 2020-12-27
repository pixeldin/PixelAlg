package tree

/*
  levelOrder 按照层级遍历一颗二叉树,
  由上至下, 每层从左至右
  思路1: 逐层遍历, 使用额外队列(先进先出)把左右子树存储
  思路2: DFS, 深度优先遍历左子树, 每进入一层层数+1, 每层存储一个单独的切片.
*/
func levelOrder(root *TreeNode) [][]int {
	return nil
}
