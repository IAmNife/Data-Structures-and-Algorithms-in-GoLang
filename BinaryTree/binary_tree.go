package main

import "fmt"

//Node represents the components of a Binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Generation of the Binary Tree
//Function will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k { //If the k value is greater than the node.Key value then move to right
		if n.Right == nil {
			n.Right = &Node{Key: k} //If the node is empty then place the value as a new node
		} else {
			n.Right.Insert(k) //Else continue the Insert function again
		}
	} else if n.Key > k { //If the k value is lesser than the node.Key value then move to left
		if n.Left == nil {
			n.Left = &Node{Key: k} //If the node is empty then place the value as a new node
		} else {
			n.Left.Insert(k) //Else continue the Insert function again
		}
	}
}

//Search will take in a key value & Return true if there is a node with that value
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(123)
	tree.Insert(234)
	tree.Insert(20)
	tree.Insert(34)
	tree.Insert(12)
	tree.Insert(334)
	tree.Insert(90)
	tree.Insert(150)
	tree.Insert(13)
	tree.Insert(24)
	tree.Insert(220)
	tree.Insert(340)
	tree.Insert(35)
	tree.Insert(74)
	tree.Insert(70)
	fmt.Println(tree.Search(7))
}
