package binary_tree

func (node *TreeNode) AreToysBalanced() bool {
	if node == nil {
		return true
	}

	leftToys := node.Left.CountToys()
	rightToys := node.Right.CountToys()

	return leftToys == rightToys
}

func (node *TreeNode) CountToys() int {
	if node == nil {
		return 0
	}

	count := 0
	if node.HasToy {
		count++
	}

	return count + node.Left.CountToys() + node.Right.CountToys()
}
