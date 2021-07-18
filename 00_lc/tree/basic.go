package tree

/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeWithThree(rootVal, leftVal, rightVal int) *TreeNode {
	left := TreeNode{
		leftVal,
		nil,
		nil,
	}
	right := TreeNode{
		rightVal,
		nil,
		nil,
	}
	root := TreeNode{
		rootVal,
		&left,
		&right,
	}
	return &root
}

func BuildTreeNode(rootVal int) *TreeNode {
	root := TreeNode{
		rootVal,
		nil,
		nil,
	}
	return &root
}

func CreateBinTreeWithArray(lt *[]int, nullNum int) *TreeNode {
	if len(*lt) == 0 {
		return nil
	}
	var node *TreeNode
	v := (*lt)[0]
	*lt = (*lt)[1:]
	if v == nullNum {
		// reach bottom
		return node
	}
	node = &TreeNode{Val: v}
	node.Left = CreateBinTreeWithArray(lt, nullNum)
	node.Right = CreateBinTreeWithArray(lt, nullNum)
	return node
}
