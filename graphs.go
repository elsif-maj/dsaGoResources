package main

import "fmt"

////////////////////
////// Vertex //////
////////////////////

// Vertex struct
type Vertex[T any] struct {
	value       T
	adjVertices []*Vertex[T]
}

func createVertex[T any](val T) *Vertex[T] {
	return &Vertex[T]{val, []*Vertex[T]{}}
}

// Add adjacent vertex method for undirected graphs
func (v *Vertex[T]) addAdjVert(adjV *Vertex[T]) {
	v.adjVertices = append(v.adjVertices, adjV)
	adjV.adjVertices = append(adjV.adjVertices, v) // is this right?
}

// Add adjacent vertex method for directed graphs
func (v *Vertex[T]) addAdjVertDirected(adjV *Vertex[T]) {
	v.adjVertices = append(v.adjVertices, adjV) // is this right?
}

///////////////////////
////// Traversal //////
///////////////////////

// Depth-First Search using iteration through collection of adjacents & recursion:
func dfsRec[T comparable](v *Vertex[T], vv map[T]bool) {
	fmt.Println(v.value)
	vv[v.value] = true
	for _, y := range v.adjVertices {
		_, ok := vv[y.value]
		if !ok {
			dfsRec(y, vv)
		}
	}
}

// Depth-First Search using a call stack:
func dfsStack[T comparable](v *Vertex[T], vv map[T]bool) {
	q := []Vertex[T]{}
	q = append(q, *v)

	for len(q) > 0 {
		cv := q[0]
		q = q[1:]
		_, ok := vv[cv.value]
		if !ok {
			fmt.Println(cv.value)
			vv[cv.value] = true
			for i := len(cv.adjVertices) - 1; i >= 0; i-- {
				ptr := cv.adjVertices[i]
				q = append([]Vertex[T]{*ptr}, q...)
			}
		}
	}
}

// Breadth-First Search:

func bfs[T comparable](v *Vertex[T]) {
	q := []*Vertex[T]{}
	m := map[T]bool{}
	q = append(q, v)
	var cv *Vertex[T]

	for len(q) > 0 {
		cv = q[0]
		q = q[1:]
		fmt.Println(q)
		_, ok := m[cv.value]
		if !ok {
			m[cv.value] = true
			q = append(q, cv.adjVertices...)
		}
	}
}

///////////////////////////
////// Program/Tests //////
///////////////////////////

func main() {
	/*
	       __ Alice __
	      /   /   \   \
	   Bob Candy Derek-Elaine
	    |    |     |
	   Fred  /    Gina
	    |   /      |
	   Helen      Irena
	*/
	// Create the above graph:
	// Begin
	alice := createVertex[string]("Alice")
	bob := createVertex[string]("Bob")
	candy := createVertex[string]("Candy")
	derek := createVertex[string]("Derek")
	elaine := createVertex[string]("Elaine")
	alice.addAdjVert(bob)
	alice.addAdjVert(candy)
	alice.addAdjVert(derek)
	alice.addAdjVert(elaine)
	fred := createVertex[string]("Fred")
	gina := createVertex[string]("Gina")
	helen := createVertex[string]("Helen")
	irena := createVertex[string]("Irena")
	bob.addAdjVert(fred)
	derek.addAdjVert(elaine)
	derek.addAdjVert(gina)
	fred.addAdjVert(helen)
	gina.addAdjVert(irena)
	candy.addAdjVert(helen)
	// end

	// Test stuff:

	// m := map[string]bool{}
	// dfsRec(alice, m)
	// dfsStack(alice, m)
	bfs(alice)
}
