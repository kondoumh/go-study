package main

import (
	"fmt"
	"github.com/yourbasic/graph"
)

func main() {
	Example_basics()
	ExampleString()
}

// Build a plain graph and visit all of its edges.
func Example_basics() {
	// Build a graph with four vertices and four undirected edges.
	// (Each of these edges are, in fact, represented by two directed
	// edges pointing in opposite directions.)
	g := graph.New(4)
	g.AddBoth(0, 1) //  0 -- 1
	g.AddBoth(0, 2) //  |    |
	g.AddBoth(2, 3) //  2 -- 3
	g.AddBoth(1, 3)

	// The vertices of all graphs in this package are numbered 0..n-1.
	// The edge iterator is a method called Visit; it calls a function
	// for each neighbor of a given vertex. Together with the Order
	// method — which returns the number of vertices in a graph — it
	// constitutes an Iterator. All algorithms in this package operate
	// on any graph implementing this interface.

	// Visit all edges of a graph.
	for v := 0; v < g.Order(); v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			// Visiting edge (v, w) of cost c.
			return
		})
	}

	// The immutable data structure created by Sort has an Iterator
	// that returns neighbors in increasing order.

	// Visit the edges in order.
	for v := 0; v < g.Order(); v++ {
		graph.Sort(g).Visit(v, func(w int, c int64) (skip bool) {
			// Visiting edge (v, w) of cost c.
			return
		})
	}

	// The return value of an iterator function is used to break
	// out of the iteration. Visit, in turn, returns a boolean
	// indicating if it was aborted.

	// Skip the iteration at the first edge (v, w) with w equal to 3.
	for v := 0; v < g.Order(); v++ {
		aborted := graph.Sort(g).Visit(v, func(w int, c int64) (skip bool) {
			fmt.Println(v, w)
			if w == 3 {
				skip = true // Aborts the call to Visit.
			}
			return
		})
		if aborted {
			break
		}
	}
	// output:
	// 0 1
	// 0 2
	// 1 0
	// 1 3
}

// Find the shortest distances from a vertex in an unweighted graph.
func ExampleBFS() {
	gm := graph.New(6)
	gm.AddBoth(0, 1) //  0--1--2
	gm.AddBoth(0, 3) //  |  |  |
	gm.AddBoth(1, 2) //  3--4  5
	gm.AddBoth(1, 4)
	gm.AddBoth(2, 5)
	gm.AddBoth(3, 4)
	g := graph.Sort(gm)

	dist := make([]int, g.Order())
	graph.BFS(g, 0, func(v, w int, _ int64) {
		fmt.Println(v, "to", w)
		dist[w] = dist[v] + 1
	})
	fmt.Println("dist:", dist)
	// Output:
	// 0 to 1
	// 0 to 3
	// 1 to 2
	// 1 to 4
	// 2 to 5
	// dist: [0 1 2 1 2 3]
}

// Print a graph.
func ExampleString() {
	g0 := graph.New(0)
	fmt.Println(g0)

	g1 := graph.New(1)
	g1.Add(0, 0)
	fmt.Println(g1)

	g4 := graph.New(4) //             8
	g4.AddBoth(0, 1)   //  0 <--> 1 <--- 2      3
	g4.AddCost(2, 1, 8)
	fmt.Println(g4)
	// Output:
	// 0 []
	// 1 [(0 0)]
	// 4 [{0 1} (2 1):8]
}


