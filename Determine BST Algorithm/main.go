package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func isBinarySearchTree(n *node) bool {
	return isBSTUtil(n, nil, nil)
}

func isBSTUtil(n *node, left *node, right *node) bool {
	if n == nil {
		return true
	}

	if left != nil && n.value < left.value {
		return false
	}

	if right != nil && n.value > right.value {
		return false
	}

	return isBSTUtil(n.left, left, n) && isBSTUtil(n.right, n, right)
}

func main() {
	root := &node{value: 4}
	root.left = &node{value: 2}
	root.right = &node{value: 5}
	root.left.left = &node{value: 1}
	root.left.right = &node{value: 3}

	if isBinarySearchTree(root) {
		fmt.Println("The tree is a binary search tree")
	} else {
		fmt.Println("The tree is not a binary search tree")
	}
}
