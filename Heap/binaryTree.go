package main

import "fmt"

// Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert method
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			// if right child is empty
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			// if left chils is empty
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search Method
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}
	// move right
	if k > n.Key {
		return n.Right.Search(k)
		// move left
	} else if k < n.Key {
		return n.Left.Search(k)
	} else {
		return true
	}
}

func main() {
	tree := &Node{Key: 100}
	for i := 100; i < 1000; i += 100 {
		// building the binary search tree
		tree.Insert(i)
	}
	// searching the element
	fmt.Println(tree.Search(800))
	fmt.Println(tree.Search(650))
}
