// In pursuit of learning pointers,impelemeted a doubly linked list in Go

package main

import "fmt"

//Define the strcuture of the doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	value  string
	after  *Node
	before *Node
}

//The data structure will have three operations: Search, Insert and Delete

//In case of Insert, when the first element is added, it will be both the head and the tail
//Elements inserted afterwards will become separate head and tails
func (d *DoublyLinkedList) InsertFirst(v string) {
	newNode := &Node{value: v}
	d.head = newNode
	d.tail = newNode
}

func (d *DoublyLinkedList) InsertAtBeginning(v string) {
	newNode := &Node{value: v}
	newNode.after = d.head
	d.head.before = newNode
	d.head = newNode
}

func (d *DoublyLinkedList) InsertAtEnd(v string) {
	newNode := &Node{value: v}
	newNode.before = d.tail
	d.tail.after = newNode
	d.tail = newNode
}

func (d DoublyLinkedList) PrintListForward() {
	for d.head != nil {
		fmt.Print(d.head.value, " -> ")
		d.head = d.head.after
	}
	fmt.Println()
}

func (d DoublyLinkedList) PrintListBackward() {
	for d.tail != nil {
		fmt.Print(d.tail.value, " -> ")
		d.tail = d.tail.before
	}
	fmt.Println()
}

func (d DoublyLinkedList) Search(v string) bool {
	nodePosition := 0
	for d.head != nil {
		if d.head.value == v {
			return true
		}
		d.head = d.head.after
		nodePosition += 1
	}
	return false
}

func (d *DoublyLinkedList) SearchAndDelete(v string) {
	//Deleting the head
	if d.head.value == v {
		// If head is being deleted
		d.head = d.head.after
		d.head.before = nil
		return
	}
	//Deleting the tail
	if d.tail.value == v {
		d.tail = d.tail.before
		d.tail.after = nil
		return
	}
	previousNode := d.head
	for previousNode.after != nil {
		if previousNode.after.value == v {
			previousNode.after.after.before = previousNode.after.after.before.before
			previousNode.after = previousNode.after.after
		} else {
			previousNode = previousNode.after
		}
	}
}

func main() {
	myDoublyLinkedList := &DoublyLinkedList{}
	myDoublyLinkedList.InsertFirst("0")
	elementsInsertForward := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	elementsInsertBackward := []string{"-1", "-2", "-3", "-4", "-5", "-6", "-7", "-8", "-9"}
	fmt.Println("Insert element at start")
	for _, elementForward := range elementsInsertForward {
		myDoublyLinkedList.InsertAtBeginning(elementForward)
		myDoublyLinkedList.PrintListForward()
	}
	fmt.Println("Insert element at end")
	for _, elementBackward := range elementsInsertBackward {
		myDoublyLinkedList.InsertAtEnd(elementBackward)
		myDoublyLinkedList.PrintListForward()
	}
	fmt.Println("Searching element")
	fmt.Println(myDoublyLinkedList.Search("3"))
	fmt.Println("Deleting the element")
	myDoublyLinkedList.SearchAndDelete("3")
	fmt.Println("Searching the element again... ")
	fmt.Println(myDoublyLinkedList.Search("3"))
	fmt.Println("Print list backwards")
	myDoublyLinkedList.PrintListBackward()
}
