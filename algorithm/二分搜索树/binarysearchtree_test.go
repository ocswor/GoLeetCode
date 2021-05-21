package algorithm

import "testing"

func TestTreeNode_Insert(t *testing.T) {

	n := TreeNode{value: -1}
	arr := []int{1, 2, 2, 4, 56, 6, 8}
	for i := 0; i < len(arr); i++ {
		node := TreeNode{value: arr[i]}
		n.Insert(&node)
	}

	t.Log(n)
}

func TestLeftPrintTreeNode(t *testing.T) {
	n := TreeNode{value: -1}
	arr := []int{1, 2, 2, 4, 56, 6, 8}
	for i := 0; i < len(arr); i++ {
		node := TreeNode{value: arr[i]}
		n.Insert(&node)
	}

	LeftPrintTreeNode(&n)

}

func TestTreeNode_Inorder(t *testing.T) {
	n := TreeNode{value: -1}
	arr := []int{1, 2, 2, 4, 56, 6, 8}
	for i := 0; i < len(arr); i++ {
		node := TreeNode{value: arr[i]}
		n.Insert(&node)
	}

	r := n.Inorder()
	t.Log(r)
}

func TestTreeNode_Preorder(t *testing.T) {
	n := TreeNode{value: -1}
	arr := []int{1, 2, 2, 4, 56, 6, 8}
	for i := 0; i < len(arr); i++ {
		node := TreeNode{value: arr[i]}
		n.Insert(&node)
	}

	r := n.Preorder()
	t.Log(r)
}
