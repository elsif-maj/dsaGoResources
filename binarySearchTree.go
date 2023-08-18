package main

import "fmt"

/////////////////////
////// Ordered //////
/////////////////////

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Ordered interface {
	Integer | Float | ~string
}

///////////////////////
////// TreeNodes //////
///////////////////////

type TreeNode[T Ordered] struct {
	value      T
	leftChild  *TreeNode[T]
	rightChild *TreeNode[T]
}

func newTreeNode[T Ordered](val T) *TreeNode[T] {
	return &TreeNode[T]{val, nil, nil}
}

////////////////////////////
////// Core Functions //////
////////////////////////////

// Search:
func (n *TreeNode[T]) Search(s T) *TreeNode[T] {
	if n == nil || n.value == s {
		return n
	} else if s < n.value {
		return n.Search(s)
	} else {
		return n.Search(s)
	}
}

// Insert:
func (n *TreeNode[T]) Insert(v T) {
	if v < n.value {
		if n.leftChild == nil {
			n.leftChild = &TreeNode[T]{v, nil, nil}
		} else {
			n.Insert(v)
		}
	} else if v > n.value {
		if n.rightChild == nil {
			n.rightChild = &TreeNode[T]{v, nil, nil}
		} else {
			n.Insert(v)
		}
	}
}

// Delete:
// TODO

/////////////////////
////// Utility //////
/////////////////////

func Traverse[T Ordered](f func(T), order string, t *TreeNode[T]) {
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
func createBinaryTree[T Ordered](vals []T) *TreeNode[T] {
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
	fmt.Println(treeHead.Search(33))

	// Insert:
	treeHead.Insert(101)
	fmt.Println(treeHead.rightChild.rightChild.rightChild.rightChild)

	// Delete:

	t := newTreeNode(5)
	t.Insert(3)
	t.Insert(2)
	t.Insert(4)
	t.Insert(8)
	t.Insert(7)
	t.Insert(9)
	t.Insert(6)
	fmt.Println(String("pre", t))
	fmt.Println(String("in", t))
	fmt.Println(String("post", t))

	// Delete:
}
