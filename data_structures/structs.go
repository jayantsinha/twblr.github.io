package data_structures

// Linked List Implementation

type Node struct{
	element int
}

func Add(elem int) {
	node := new(Node)
	node.element = elem
	return
}

func Search(elem int) *Node {
	f:=Node{elem}
	return &f
}
