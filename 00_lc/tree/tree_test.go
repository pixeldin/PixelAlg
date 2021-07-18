package tree

import "testing"

func buildLocalTree() *TreeNode {
	a := BuildTreeWithThree(1, 2, 3)
	b := BuildTreeWithThree(2, 4, 5)
	c := BuildTreeWithThree(3, 6, 7)
	d := BuildTreeWithThree(4, 8, 9)
	e := BuildTreeWithThree(5, 10, 11)
	f := BuildTreeWithThree(6, 12, 13)
	g := BuildTreeWithThree(7, 15, 16)
	a.Left = b
	a.Right = c

	b.Left = d
	b.Right = e

	c.Left = f
	c.Right = g

	return a
}

func TestLevelOrder(t *testing.T) {
	tree := buildLocalTree()
	//levelOrder(tree)
	zigzagLevelOrder(tree)
}

func TestNumTree(t *testing.T) {
	numTrees(3)
}

func TestIsBlancedTree(t *testing.T) {
	//tr := CreateBinTreeWithArray(&[]int{1,2,3,4,-1,-1,4,-1,-1,3,-1,-1,2,-1,-1})
	//tr := CreateBinTreeWithArray(&[]int{1,-1,2,-1,3,-1,-1})
	//tr := CreateBinTreeWithArray(&[]int{1,2,4,8,-1,-1,-1,5,-1,-1,3,6,-1,-1,-1})
	tr := CreateBinTreeWithArray(&[]int{1, 0, 2, 0, 3, 0, 0}, 0)
	t.Log(isBalanced(tr))
}
