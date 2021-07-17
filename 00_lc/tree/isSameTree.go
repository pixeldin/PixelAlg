package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return checkVal(p, q)
}

func checkVal(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if q == nil && p == nil {
		return true
	}
	if q != nil && p != nil {
		if q.Val != p.Val {
			return false
		}
	}
	return checkVal(p.Left, q.Left) && checkVal(p.Right, q.Right)
}

func TestisSameTree() bool {
	three := BuildTreeWithThree(1, 3, 3)
	three2 := BuildTreeWithThree(1, 2, 3)
	same := isSameTree(three, three2)
	return same
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return checkVal(p, q)
}
