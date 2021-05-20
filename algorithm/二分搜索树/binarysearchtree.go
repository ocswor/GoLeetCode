package algorithm

type TreeNode struct {
	Value int
	left  *TreeNode
	right *TreeNode
	count int
}

func (n *TreeNode) Insert(node *TreeNode, value int) {

	if node == nil {
		n.count++
		n.Value = value
	}
}
