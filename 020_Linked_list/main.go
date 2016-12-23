package main

import (
	"fmt"
	"time"
)

// Implement a linked list (single linked)
// list.head() - moves current node to head node
// list.tail() - moves current node to the last node
// list.delete() - deletes current node

type NodeData struct {
	data int
	name string
}

type Node struct {
	data NodeData
	next *Node
}

type LinkedList struct {
	head    *Node
	current *Node
	tail    *Node
}

// list.append() - adds node to the end
func (ll *LinkedList) append(n Node) {
	if ll.head == nil {
		ll.head = &n
		ll.current = ll.head
		ll.tail = ll.head
	} else {
		ll.tail.next = &n
		ll.tail = &n
	}
}

// list.next() - moves current to the next node
func (ll *LinkedList) next() bool {
	if ll.current.next != nil {
		ll.current = ll.current.next
		return true
	}
	return false
}

// We can not insert instead of current - because we don't know the previous one
func (ll *LinkedList) insertAfterCurrent(n Node) {
	n.next = ll.current.next
	ll.current.next = &n
}

// list.getCurrent() - returns current node
func (ll *LinkedList) getCurrent() *Node {
	return ll.current
}

func (ll *LinkedList) print() {
	printNodesFrom(ll.head)
	fmt.Println()
}

func printNodesFrom(n *Node) {
	fmt.Printf("%v", n.data)
	if n.next != nil {
		fmt.Printf("->")
		printNodesFrom(n.next)
	}

}

func main() {

	sll := LinkedList{nil, nil, nil}
	node := Node{NodeData{4, ""}, nil}
	node2 := Node{NodeData{5, ""}, nil}
	sll.append(node)
	fmt.Printf("Current: %v\n", sll.getCurrent())
	sll.append(node2)
	sll.insertAfterCurrent(Node{NodeData{99, "Wsadzony"}, nil})
	sll.append(Node{NodeData{1, "Kuba"}, nil})
	fmt.Printf("Current: %v\n", sll.getCurrent())
	sll.next()
	sll.next()
	sll.next()
	sll.insertAfterCurrent(Node{NodeData{99, "Durgi wsadzony"}, nil})
	sll.print()

	fmt.Printf("%v", sll.head)

	var start time.Time = time.Now()

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
