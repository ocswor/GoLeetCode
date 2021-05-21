package algorithm

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

var count int

func (n *TreeNode) Insert(node *TreeNode) {

	if n.value == -1 {
		n.value = node.value
		count++
		return
	}

	if n.value > node.value {
		if n.left == nil {
			n.left = node
		} else {
			n.left.Insert(node)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.Insert(node)
		}

	}

}

func LeftPrintTreeNode(node *TreeNode) {

	if node.left != nil {
		LeftPrintTreeNode(node.left)
	}
	println(node.value)
	if node.right != nil {
		LeftPrintTreeNode(node.right)
	}

}

// Inorder returns inorder traversal of TreeNode
func (t *TreeNode) Inorder() []int {
	if t == nil {
		return nil
	}

	return append(
		append(
			t.left.Inorder(),
			t.value,
		),
		t.right.Inorder()...,
	)
}

func (t *TreeNode) Preorder() []int {
	if t == nil {
		return nil
	}

	return append(
		append(
			[]int{t.value},
			t.left.Preorder()...,
		),
		t.right.Preorder()...,
	)
}
