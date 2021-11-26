/*
	Directed Graph using Adjacency list
*/

package main

import "fmt"

// Graph is defined as a slice of Vertices
type Graph struct {
	Vertices []*Vertex
}

// Each Vertex contains a value and an 'adjacency' list signifying the vertex neighbour
type Vertex struct {
	Key      int
	adjacent []*Vertex
}

func (g *Graph) Print() {
	for _, vertex := range g.Vertices {
		fmt.Printf("\nVertex %v : ", vertex.Key)
		for _, v := range vertex.adjacent {
			fmt.Printf(" %v ", v.Key)
		}
	}
	fmt.Println()
}

func (g *Graph) AddEdge(from, to int) error {
	if g.Contains(from) && g.Contains(to) {
		for _, vertex := range g.Vertices {
			if from == vertex.Key {
				for _, v := range vertex.adjacent {
					if v.Key == to {
						return fmt.Errorf("Edge from %v -> %v already exists", from, to)
					}
				}
				vertex.adjacent = append(vertex.adjacent, &Vertex{Key: to})
			}
		}
		return nil
	} else {
		if !g.Contains(from) {
			return fmt.Errorf("Vertex %v is not in graph", from)
		} else {
			return fmt.Errorf("Vertex %v is not in graph", to)
		}
	}
}

func (g *Graph) AddVertex(k int) error {
	if !g.Contains(k) {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
		return nil
	} else {
		return fmt.Errorf("Vertex %v not added as it is an existing key", k)
	}
}

func (g Graph) Contains(k int) bool {
	for _, vertex := range g.Vertices {
		if k == vertex.Key {
			return true
		}
	}
	return false
}

func main() {
	myGraph := &Graph{}
	for i := 0; i < 5; i++ {
		if err := myGraph.AddVertex(i); err != nil {
			fmt.Println(err)
		}
	}
	if err := myGraph.AddEdge(1, 2); err != nil {
		fmt.Println(err)
	}
	if err := myGraph.AddEdge(1, 2); err != nil {
		fmt.Println(err)
	}
	if err := myGraph.AddEdge(100, 2); err != nil {
		fmt.Println(err)
	}
	myGraph.Print()
}
