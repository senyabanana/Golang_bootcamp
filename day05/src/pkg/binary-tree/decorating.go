package binary_tree

func (node *TreeNode) UnrollGarland() []bool {
	if node == nil {
		return nil
	}

	var result []bool
	queue := []*TreeNode{node}
	level := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := make([]bool, levelSize)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			index := i
			if level%2 != 0 {
				index = levelSize - i - 1
			}
			levelValues[index] = current.HasToy

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		result = append(result, levelValues...)
		level++
	}

	return result
}
