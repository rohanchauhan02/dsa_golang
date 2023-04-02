package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

//[STACK] Push will add a value at the end

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// [STACK] Pop will remove a value from end and return the value

func (s *Stack) Pop() int {
	remove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return remove
}


// Queue represents a queue that holds a slice

type Queue struct{
	items []int
}

// [QUEUE] Enqueue item at the last
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//  [QUEUE] Dequeue will remove a value from satat and return the value
func (q *Queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[1:]
	return remove
}

func main() {

	// STACK
	myStack := Stack{
		items: []int{1, 2, 3, 4, 5, 6},
	}
	fmt.Printf("STACK :%v\n",myStack)
	myStack.Push(3)
	fmt.Printf("STACK :%v\n",myStack)
	fmt.Printf("POP Item: %v ",myStack.Pop())
	fmt.Printf("STACK :%v\n",myStack)

	//QUEUE
	myQueue := Queue{
		items: []int{1, 2, 3, 4, 5, 6},
	}
	fmt.Println(myQueue)
	myQueue.Enqueue(3)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue)

}
