package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (ll *linkedlist) Add(n *node) {
	second := ll.head
	ll.head = n
	ll.head.next = second
	ll.length ++
}

func (ll linkedlist) PrintListData() {
	toPrint := ll.head
	for ll.length != 0 {
		fmt.Printf("%d <- ", toPrint.data)
		toPrint = toPrint.next
		ll.length --
	}
	fmt.Println()
}

// Delete
func (ll *linkedlist) DeleteWithValue(value int){
}





func main() {
	//Linkedlist

	myLinkedList := linkedlist{}

	myLinkedList.Add(&node{data: 4, next: nil})
	myLinkedList.Add(&node{data: 1, next: nil})
	myLinkedList.Add(&node{data: 3, next: nil})
	myLinkedList.Add(&node{data: 2, next: nil})

	myLinkedList.PrintListData()

}
