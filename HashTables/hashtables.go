/*
Based on example of Junmin Lee (YouTube)

Implemented Hash Table as an array of linked list
Each element of the array is a linked list
If in case there is a hash value collision, new element is prepended to the linked list at the position of hash collision
*/

package main

import "fmt"

//Define Hash table size
const HashTableSize = 10

//Define hash table as an array
type HashTable struct {
	HashTableEntry [HashTableSize]*Bucket
}

//Define each element as a linked list
type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

//Define Insert, Search, Delete for HashTable
func (h *HashTable) Insert(key string) {
	hashValue := Hash(key)
	h.HashTableEntry[hashValue].InsertBucket(key)
}

func (h *HashTable) Search(key string) bool {
	hashValue := Hash(key)
	return h.HashTableEntry[hashValue].SearchBucket(key)
}

func (h *HashTable) Delete(key string) {
	hashValue := Hash(key)
	h.HashTableEntry[hashValue].DeleteBucket(key)
}

//Define InsertBucket, SearchBucket, DeleteBucket to operate on the Bucket (linked list level)
//For insert, I will Insert the new node to the bucket linked list
func (b *Bucket) InsertBucket(k string) {
	//If only SearchBucket returns false do we initiate insert
	if !b.SearchBucket(k) {
		NewNode := &BucketNode{key: k}
		NewNode.next = b.head //NewNode's next should point to the initial first node
		b.head = NewNode      //Change the head of linked list to NewNode
	} else {
		fmt.Println("The value already exists")
	}
}

//For search, I will search through each entry in the linked list sequentially; if value not found, return false; else true
func (b *Bucket) SearchBucket(k string) bool {
	//Start searching from the head
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//For delete, I search through each entry in the linked list, and remove if there is a match
func (b *Bucket) DeleteBucket(k string) {
	//Two cases here;
	//case 1: If value to delete is actually in the head node
	if b.head.key == k {
		//make the linked list head point to the second node
		b.head = b.head.next
		//return after this case is evaluated
		return
	}
	//case 2: The value resides not in the head node
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}

}

//Define a Hash calculating function
func Hash(key string) int {
	hashval := 0
	for _, v := range key {
		hashval += int(v)
	}
	return hashval % HashTableSize
}

//Initialize the hash table array
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.HashTableEntry {
		result.HashTableEntry[i] = &Bucket{}
	}
	return result
}

func main() {
	//Initialize the hash table
	myHashTable := Init()

	//Define entries to insert to Hash table
	entries := []string{
		"DYLAN",
		"BOXER",
		"HAMMER",
		"CRABS",
		"DYLAN", //insert duplicate value to check Insert(k) duplicate entry handling logic
	}

	for _, value := range entries {
		//Insert entries into hash table
		myHashTable.Insert(value)
	}

	//search for an entry which was never inserted, should return false
	fmt.Println(myHashTable.Search("DEXTER"))
	//search for an entry which was inserted, shoudl return true
	fmt.Println(myHashTable.Search("DYLAN"))
	//delete an existing entry
	myHashTable.Delete("DYLAN")
	//search for the deleted entry, should return false
	fmt.Println(myHashTable.Search("DYLAN"))

}
