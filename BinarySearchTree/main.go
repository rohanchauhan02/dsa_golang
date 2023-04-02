package main

import (
	"fmt"
	"strconv"
)

// Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert

func (n *Node) Insert(k int) {
	if n.Key < k {
		//Move Right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move Left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Print
func (n *Node) PrintTree() {
	if n == nil {
		return
	}
	s := ""
	if n.Left == nil {
		s += "."
	} else {
		s += strconv.Itoa(n.Left.Key)
	}
	s += " -> " + strconv.Itoa(n.Key) + " <- "
	if n.Right == nil {
		s += "."
	} else {
		s += strconv.Itoa(n.Right.Key)
	}
	fmt.Println(s)
	n.Left.PrintTree()
	n.Right.PrintTree()
}

//Search

func main() {
	tree := &Node{Key: 100}

	tree.Insert(50)
	tree.Insert(101)
	tree.Insert(77)
	tree.Insert(105)
	tree.Insert(33)
	tree.Insert(155)
	tree.Insert(1)
	tree.PrintTree()

}
