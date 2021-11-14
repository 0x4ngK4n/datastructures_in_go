package main

import "fmt"

//Stack

//Define a stack
//So stack is a collection of Last in, First out elements
type Stack struct {
	stackElements []int
}

//Define what operations can be performed over Stack, such as -> Push and Pop
//Push is to append to the slice
func (s *Stack) Push(i int) {
	s.stackElements = append(s.stackElements, i)
}

//Pop is to remove the last element of the slice
func (s *Stack) Pop() int {
	stackLen := len(s.stackElements)
	if stackLen != 0 {
		s.stackElements = s.stackElements[:stackLen-1]
		return stackLen
	}
	return 0
}

//Queue

//Define a Queue
//So queue is a collection of Last in , last out
type Queue struct {
	queueElements []int
}

//Ditto as Stack Push()
func (q *Queue) EnQueue(i int) {
	q.queueElements = append(q.queueElements, i)
}

//In DeQueue, we remove the first element of the slice
func (q *Queue) DeQueue() int {
	queueLength := len(q.queueElements) - 1
	if queueLength > 0 {
		for i := 0; i < queueLength; i++ {
			q.queueElements[i] = q.queueElements[i+1]
		}
		q.queueElements = q.queueElements[:queueLength]
		return queueLength
	} else {
		return 0
	}
}

func main() {
	//Items for operation on Stack and Queue
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Stack")
	myStack := &Stack{}
	for _, v := range items {
		myStack.Push(v)
		fmt.Println(myStack.stackElements)
	}
	for myStack.Pop() > 1 {
		fmt.Println(myStack.stackElements)
	}
	fmt.Println("===============================")
	fmt.Println("Queue")
	myQueue := &Queue{}
	for _, v := range items {
		myQueue.EnQueue(v)
		fmt.Println(myQueue.queueElements)
	}
	for myQueue.DeQueue() > 0 {
		fmt.Println(myQueue.queueElements)
	}

}
