package main

import (
	"fmt"

	tree "binary-tree/pkg/binary-tree"
)

func main() {
	root := tree.NewTreeNode(true)
	root.Left = tree.NewTreeNode(true)
	root.Right = tree.NewTreeNode(false)
	root.Right.Right = tree.NewTreeNode(true)
	root.Right.Left = tree.NewTreeNode(true)
	root.Left.Right = tree.NewTreeNode(false)
	root.Left.Left = tree.NewTreeNode(true)

	/*
		    1
		   /  \
		  1     0
		 / \   / \
		1   0 1   1
	*/
	fmt.Println(root.UnrollGarland())
}
