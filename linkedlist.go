package main

import "fmt"

/////////////////////////////////////////////////
//////////////// Node Basics: ///////////////////
/////////////////////////////////////////////////

type Node[T comparable] struct {
	data T
	next *Node[T]
}

// Node Factory Function:
func CreateNode[T comparable](data T, next *Node[T]) *Node[T] {
	return &Node[T]{data, next}
}

////////////////////////////////////////////////////////
//////////////// Linked List Basics: ///////////////////
////////////////////////////////////////////////////////

// LinkedList Struct:
type LinkedList[T comparable] struct {
	head *Node[T]
}

// LinkedList Factory Function:
func CreateLinkedList[T comparable](head *Node[T]) *LinkedList[T] {
	return &LinkedList[T]{head}
}

/////////////////////////////////////////////////////////
//////////////// Linked List Methods: ///////////////////
/////////////////////////////////////////////////////////

// Read the data value at an index.
func (ll *LinkedList[T]) Read(index int) interface{} {
	cn := ll.head
	for i := 0; i < index; i++ {
		if cn.next == nil {
			return nil
		}
		cn = cn.next
	}
	return cn.data
}

// Search for the index of a data value. The reason for choosing an int pointer
// rather than just an int as a return value is so that we can return nil.
func (ll *LinkedList[T]) Search(val T) int {
	cn := ll.head
	for i := 0; ; i++ {
		if cn == nil {
			return -1
		}
		if cn.data == val {
			return i
		} else {
			cn = cn.next
		}
	}
}

// Insert a new Node into a Linked List at some index position
func (ll *LinkedList[T]) Insert(idx int, val T) {
	// New node:
	nn := CreateNode(val, nil)

	// If inserting at the beginning of a list:
	if idx == 0 {
		nn.next = ll.head
		ll.head = nn
		return
	}

	// "current node" and "current index" for finding position in list where insertion will take place.
	// Insertion position is one position before the index that will be occupied by the newly inserted
	// node. e.g. node1 -> !node2! -> node3 for idx argument "2"
	cn := ll.head
	ci := 0
	for ; ci < idx-1; ci++ {
		if cn == nil {
			return
		}
		cn = cn.next
	}

	nn.next = cn.next
	cn.next = nn
}

// Delete a node from the list:
func (ll *LinkedList[T]) Delete(idx int) {
	if idx == 0 {
		ll.head = ll.head.next
		return
	}

	cn := ll.head
	ci := 0
	for ; ci < idx-1; ci++ {
		if cn == nil {
			return
		}
		cn = cn.next
	}

	cn.next = cn.next.next
}

// Walk a linked list given pointer to LinkedList struct:
func (ll *LinkedList[T]) Walk() {
	cn := ll.head
	s := ""

	for idx := 0; cn != nil; idx++ {
		s += fmt.Sprintf("%v", cn.data)
		if cn.next != nil {
			s += " -> "
		}

		cn = cn.next
	}
	fmt.Println(s)
}

// Walk a linked list given the head node:
// TODO: make this

/////////////////////////////////////////////////////////
//////////////// Testing & App Logic: ///////////////////
/////////////////////////////////////////////////////////

func main() {
	// Node setup: node1 -> node2 -> node3 -> node4 AND created Linked List called "ll"
	node1 := CreateNode("hello", nil)
	node2 := CreateNode("there", nil)
	node3 := CreateNode("my", nil)
	node4 := CreateNode("friend", nil)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	ll := CreateLinkedList[string](node1)
	// Run tests/logic:

	// Read:
	fmt.Println(ll.Read(3) == "friend") // Test our "Read" function for returning a node's data value by index
	fmt.Println(ll.Read(7) == nil)      // Test our "Read" function for not finding an index

	// Search:
	fmt.Println(ll.Search("friend") == 3) // Test our "Search" function for returning the index of a data value
	fmt.Println(ll.Search("stuff") == -1) // Test our "Search" function for returning -1 if it doesn't find a data value

	// Insert:
	ll.Insert(3, "dear")
	ll.Walk()

	// Delete:
	ll.Delete(3)
	ll.Walk()
}
