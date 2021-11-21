/*
Circular queue implementation in Go
*/

package main

import (
	"errors"
	"fmt"
)

// Define circular queue size
var queueSize = 5

// I need two trackers, start and end.
type CircularQueue struct {
	start         int
	end           int
	queueElements []int
}

// EnQueue (Insert) in circular queue
func (c *CircularQueue) EnQueue(element int) error {
	if (c.end - c.start) < queueSize {
		c.queueElements = append(c.queueElements, element)
		c.end += 1
		return nil
	} else {
		return errors.New("queue is full")
	}
}

// DeQueue (Delete) from circular queue
func (c *CircularQueue) DeQueue() error {
	if c.start <= c.end {
		c.start += 1
		return nil
	} else {
		return errors.New("queue is empty")
	}
}

func (c CircularQueue) Print() {
	for i := c.start; i < c.end; i++ {
		fmt.Print(c.queueElements[i], "|")
	}
	fmt.Println()
}

func main() {
	myCircularQueue := &CircularQueue{start: 0, end: 0}
	toInsertItems := []int{1, 2, 3, 4, 5}
	for _, value := range toInsertItems {
		myCircularQueue.EnQueue(value)
	}
	myCircularQueue.Print()
	myCircularQueue.DeQueue()
	myCircularQueue.DeQueue()
	myCircularQueue.DeQueue()
	myCircularQueue.DeQueue()
	myCircularQueue.Print()
	myCircularQueue.EnQueue(6)
	myCircularQueue.EnQueue(7)
	myCircularQueue.Print()
	myCircularQueue.DeQueue()
	myCircularQueue.Print()
}
