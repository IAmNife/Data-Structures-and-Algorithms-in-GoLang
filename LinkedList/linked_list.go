package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	tail int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.tail++
}

func (l LinkedList) printListData() {
	toPrint := l.Head
	for l.tail != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.tail--
	}
	fmt.Println("\n")
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.tail == 0 { //If empty list
		return
	}

	if l.Head.Data == value { //If the head is deleted, the head value is shifted to the next value
		l.Head = l.Head.Next
		l.tail--
		return
	}

	//If the value to be deleted is not in the list
	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.tail--
}

func main() {
	myLiist := LinkedList{}
	node1 := &Node{Data: 48}
	node2 := &Node{Data: 18}
	node3 := &Node{Data: 16}
	node4 := &Node{Data: 45}
	node5 := &Node{Data: 12}
	node6 := &Node{Data: 28}
	myLiist.prepend(node1)
	myLiist.prepend(node2)
	myLiist.prepend(node3)
	myLiist.prepend(node4)
	myLiist.prepend(node5)
	myLiist.prepend(node6)
	myLiist.printListData()
	myLiist.deleteWithValue(28)
	myLiist.printListData()
}
