package binary_tree

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(toy bool) *TreeNode {
	return &TreeNode{
		HasToy: toy,
		Left:   nil,
		Right:  nil,
	}
}
