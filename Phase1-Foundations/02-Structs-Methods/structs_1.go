package main

import "fmt"

type linkedList struct {
	HeadNode *Node
	TailNode *Node
	// NextNode *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) SetValue(newValue int) {
	n.Value = newValue
}

func (l *linkedList) InsertNode(value int) {
	if l.HeadNode == nil {
		l.HeadNode = &Node{Value: value}
		l.TailNode = &Node{Value: value}
	} else {
		node := l.HeadNode
		for node.Next != nil {
			node = node.Next
		}
		node.Next = &Node{Value: value}
	}
}

func (l *linkedList) printList() {
	head := l.HeadNode
	for head != nil {
		fmt.Printf("%d \n", head.Value)
		head = head.Next
	}
}

func runLinked() {
	list := linkedList{}
	list.InsertNode(10)
	list.InsertNode(29)
	list.InsertNode(34)
	list.InsertNode(12)
	list.InsertNode(54)
	list.printList()
}

func main() {
	runLinked()
}
