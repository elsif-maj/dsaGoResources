package main

import "fmt"

///////////////////////
////// TreeNodes //////
///////////////////////

type TreeNode[T any] struct {
	value      T
	leftChild  *TreeNode[T]
	rightChild *TreeNode[T]
}

/////////////////////
////// Utility //////
/////////////////////

// Creates a binary tree given a slice of node values listed in 'level-order'
func createBinaryTree[T any](vals []T) *TreeNode[T] {
	nn := len(vals)
	qLike := []*TreeNode[T]{}

	// Push to our Queue-like the correct number ('nn') of nodes, each one with its value
	// provided from the vals argument
	for i := 0; i < nn; i++ {
		thisVal := vals[i]
		qLike = append(qLike, &TreeNode[T]{thisVal, nil, nil})
	}

	// Connect all nodes based on BFS-style exhaustive exploration/connection
	// of nodes at a given depth for a binary tree
	for i := 0; i < len(qLike); i++ {
		tn := qLike[i]
		d2left := i + 1

		if i+d2left < len(qLike) {
			tn.leftChild = qLike[i+d2left]
		}
		if i+(d2left+1) < len(qLike) {
			tn.rightChild = qLike[i+d2left+1]
		}
	}

	return qLike[0]
}

func main() {
	/* Create the following tree:

	        50
	       /   \
	    25	    75
	    / \     /  \
	  10   33  56   89
	 / \  / \  / \   / \
	4 11 30 40 52 61 82 95

	*/
	treeVals := []int{50, 25, 75, 10, 33, 56, 89, 4, 11, 30, 40, 52, 61, 82, 95}
	fmt.Println(createBinaryTree[int](treeVals))
}
