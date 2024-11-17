package main

import (
	"fmt"

	tree "binary-tree/pkg/binary-tree"
)

func main() {
	root1 := tree.NewTreeNode(false)
	root1.Left = tree.NewTreeNode(false)
	root1.Right = tree.NewTreeNode(true)
	root1.Left.Left = tree.NewTreeNode(false)
	root1.Left.Right = tree.NewTreeNode(true)

	/*
		    0
		   / \
		  0   1
		 / \
		0   1
	*/
	fmt.Println(root1.AreToysBalanced())

	root2 := tree.NewTreeNode(true)
	root2.Left = tree.NewTreeNode(true)
	root2.Right = tree.NewTreeNode(false)
	root2.Left.Left = tree.NewTreeNode(true)
	root2.Left.Right = tree.NewTreeNode(false)
	root2.Right.Left = tree.NewTreeNode(true)
	root2.Right.Right = tree.NewTreeNode(true)

	/*
		    1
		   /  \
		  1     0
		 / \   / \
		1   0 1   1

	*/
	fmt.Println(root2.AreToysBalanced())

	root3 := tree.NewTreeNode(true)
	root3.Left = tree.NewTreeNode(true)
	root3.Right = tree.NewTreeNode(false)

	/*
		  1
		 / \
		1   0
	*/
	fmt.Println(root3.AreToysBalanced())

	root4 := tree.NewTreeNode(false)
	root4.Left = tree.NewTreeNode(true)
	root4.Right = tree.NewTreeNode(false)
	root4.Left.Right = tree.NewTreeNode(true)
	root4.Right.Right = tree.NewTreeNode(true)

	/*
		  0
		 / \
		1   0
		 \   \
		  1   1
	*/
	fmt.Println(root4.AreToysBalanced())
}
