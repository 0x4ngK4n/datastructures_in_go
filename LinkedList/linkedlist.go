package main

import "fmt"

//Define the linked list structure
type LinkedList struct {
	head *Node
}

//Define individual nodes of a linked list
//In this linkedlist, on each node, I store a string
type Node struct {
	value string
	next  *Node
}

//Initialize the linked list
func Init() *LinkedList {
	myList := &LinkedList{}
	return myList
}

//Initialize new node
func initNode(v string) *Node {
	node := &Node{value: v}
	return node
}

//Operations on linked list such as Search, Insert, Print and Delete
//In search, I will try to iterate each node of the linked list to search for a value
func (l *LinkedList) Search(v string) bool {
	//start searching from the head of the linked list
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == v {
			return true
		}
		//make current node the next node
		currentNode = currentNode.next
	}
	return false
}

//When inserting to linked list, I am inserting new items at the head of the linked list
func (l *LinkedList) Insert(v string) {
	newNode := &Node{value: v}
	newNode.next = l.head
	l.head = newNode
}

//Similar to search, print value inside each node of the linked list
func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf(currentNode.value + " -> ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *LinkedList) Delete(v string) {
	//In case of delete, we have two cases
	//Case1: Is it the first node which has to be deleted
	if l.head.value == v {
		l.head = l.head.next
		return
	}
	//Case2: Any other element other than the first node
	previousNode := l.head
	for previousNode.next != nil {
		//If previousNode.next (current node) value is the one we want to delete
		if previousNode.next.value == v {
			//If in case the node we want to delete is the last node in linked list
			if previousNode.next.next == nil {
				previousNode.next = nil
				return
			}
			//Make previousNode point to the node after the current node (previousNode.next.next)
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func main() {
	myList := Init()
	nodesToMake := []string{
		"STOINIS",
		"WADE",
		"FINCH",
		"MAXWELL",
		"SMITH",
		"WARNER",
	}
	for _, v := range nodesToMake {
		myList.Insert(v)
	}
	myList.Print()
	myList.Delete("STOINIS")
	myList.Print()
	fmt.Println(myList.Search("STOINIS"))
}
