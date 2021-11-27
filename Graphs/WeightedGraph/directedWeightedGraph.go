package main

import "fmt"

// Define graph structure using adjacency matrix
type Graph struct {
	Vertices Matrix
	Size     int
}

type Matrix struct {
	Elements [][]int
}

// Initialize the adjacency matrix
func (g *Graph) Init() *Graph {
	tmp := make([][]int, g.Size)
	for i := range tmp {
		tmp[i] = make([]int, g.Size)
	}
	g.Vertices.Elements = tmp
	return g
}

// Print the adjacency matrix
func (g Graph) Print() {
	for i := range g.Vertices.Elements {
		fmt.Println(g.Vertices.Elements[i])
	}
	fmt.Println()
}

// Add Edge to the adjacency matrix
func (g *Graph) AddEdge(from, to, weight int) error {
	if g.Size > from && g.Size > to {
		if g.Vertices.Elements[from][to] == 0 {
			g.Vertices.Elements[from][to] = weight
			return nil
		} else {
			return fmt.Errorf("Edge %v -> %v already exists", from, to)
		}
	} else {
		return fmt.Errorf("Edge %v -> %v is not in range", from, to)
	}
}

// Remove Edge to the adjacency matrix
func (g *Graph) RemoveEdge(from, to int) error {
	if g.Size > from && g.Size > to {
		if g.Vertices.Elements[from][to] != 0 {
			g.Vertices.Elements[from][to] = 0
			return nil
		} else {
			return fmt.Errorf("Edge %v -> %v does not exists", from, to)
		}
	} else {
		return fmt.Errorf("Edge %v -> %v is not in range", from, to)
	}
}

func main() {
	myGraph := &Graph{}
	myGraph.Size = 5
	myGraph.Init()
	// Print the initialized graph
	myGraph.Print()
	if err := myGraph.AddEdge(1, 1, 1); err != nil {
		fmt.Println(err)
	}
	// Print after adding the edge
	myGraph.Print()
	if err := myGraph.AddEdge(1, 2, 5); err != nil {
		fmt.Println(err)
	}
	if err := myGraph.RemoveEdge(1, 1); err != nil {
		fmt.Println(err)
	}
	// Final print after adding the new edge and removing the previous edge
	myGraph.Print()
}
