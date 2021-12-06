package main

import "fmt"

// MaxHeap struct
type MaxHeap struct {
	array []int
}

// Insert
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

//Extract
func (h *MaxHeap) Extract() int {

	if len(h.array) == 0 {
		println("Cannot extract from empty array")
		return -1
	}

	extracted := h.array[0]
	lastElement := len(h.array) - 1
	h.array[0] = h.array[lastElement]
	h.array = h.array[:lastElement]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	// loop while index has at least one child
	lastIndex := len(h.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when only one child (i.e only left child)
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < len(buildHeap); i++ {
		m.Extract()
		fmt.Println(m)
	}
}
