package _4_tree

import "testing"

func TestCreateBinTree(t *testing.T) {
	//ls := []int{1, 2, 3, -1, -1, -1, -1}//}, 4, 5, 6, 7, 8}
	ls := []int{1, 2, 3, -1, -1, 4, -1, -1, 5, -1, 6}
	tree := CreateBinTree(&ls)
	t.Log(tree)
}

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.InOrderTraverse()
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PreOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse()
}

func TestBinaryTree_PostOrderTraverse2(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse2()
}
