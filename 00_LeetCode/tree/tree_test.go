package tree

import "testing"

func buildLocalTree() *TreeNode {
	a := BuildTreeWithThree(1, 2, 3)
	b := BuildTreeWithThree(2, 4, 5)
	c := BuildTreeWithThree(3, 6, 7)
	d := BuildTreeWithThree(4, 8, 9)
	e := BuildTreeWithThree(5, 10, 11)
	f := BuildTreeWithThree(6, 12, 13)
	g := BuildTreeWithThree(7, 12, 13)
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
	levelOrder(tree)
}
