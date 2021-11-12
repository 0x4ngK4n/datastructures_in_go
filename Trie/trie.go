//Code from tutorials from Junmin Lee.
//Had fun understanding and writing these code.

package main

import "fmt"

const AlphabetSize = 26

//Define Trie struct

//Define Trie Node
type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool
}

//Define Trie Root
type Trie struct {
	Root *Node
}

//Initialize the Trie
func InitTrie() *Trie {
	result := &Trie{Root: &Node{}}
	return result
}

//Insert in Trie
func (t *Trie) Insert(w string) {
	wordLen := len(w)
	currentNode := t.Root
	for i := 0; i < wordLen; i++ {
		checkIndex := w[i] - 'a'
		if currentNode.Children[checkIndex] == nil {
			currentNode.Children[checkIndex] = &Node{}
		}
		currentNode = currentNode.Children[checkIndex]
	}
	currentNode.IsEnd = true
}

//Search in Trie
func (t *Trie) Search(w string) bool {
	wordLen := len(w)
	currentNode := t.Root
	for i := 0; i < wordLen; i++ {
		checkIndex := w[i] - 'a'
		if currentNode.Children[checkIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[checkIndex]
	}
	if currentNode.IsEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	mySlice := []string{"aragorn",
		"aragog",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo"}
	for _, word := range mySlice {
		myTrie.Insert(word)
	}
	fmt.Println(myTrie.Search("oreo"))
}

