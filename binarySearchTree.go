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

func newTreeNode[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{val, nil, nil}
}

////////////////////////////
////// Core Functions //////
////////////////////////////

// Search:
func Search[T int](s T, n *TreeNode[T]) *TreeNode[T] {
	if n == nil || n.value == s {
		return n
	} else if s < n.value {
		return Search(s, n.leftChild)
	} else {
		return Search(s, n.rightChild)
	}
}

// Insert:
func Insert[T int](v T, n *TreeNode[T]) {
	if v < n.value {
		if n.leftChild == nil {
			n.leftChild = &TreeNode[T]{v, nil, nil}
		} else {
			Insert(v, n.leftChild)
		}
	} else if v > n.value {
		if n.rightChild == nil {
			n.rightChild = &TreeNode[T]{v, nil, nil}
		} else {
			Insert(v, n.rightChild)
		}
	}
}

// Delete:
// TODO

/////////////////////
////// Utility //////
/////////////////////

func Traverse(f func(int), order string, t *TreeNode[int]) {
	if t == nil {
		return
	}
	if order == "pre" {
		f(t.value)
		Traverse(f, order, t.leftChild)
		Traverse(f, order, t.rightChild)
	} else if order == "in" {
		Traverse(f, order, t.leftChild)
		f(t.value)
		Traverse(f, order, t.rightChild)
	} else if order == "post" {
		Traverse(f, order, t.leftChild)
		Traverse(f, order, t.rightChild)
		f(t.value)
	}
}

func String(order string, t *TreeNode[int]) string {
	s := ""
	f := func(val int) {
		s += fmt.Sprintf("%d ", val)
	}
	Traverse(f, order, t)
	return s
}

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
	treeHead := createBinaryTree[int](treeVals)

	// Search:
	fmt.Println(Search[int](33, treeHead))

	// Insert:
	Insert[int](101, treeHead)
	fmt.Println(treeHead.rightChild.rightChild.rightChild.rightChild)

	// Delete:
}
