package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func addNode(node *Node, v int) int {
	if node == nil {
		return 0
	}
	if v == node.Value {
		fmt.Println("already exist")
		return -1
	}
	if node.Next == nil {
		node.Next = &Node{v, node, nil}
		return -2
	}
	return addNode(node.Next, v)
}

func traverse(node *Node) {
	if node == nil {
		fmt.Println("empty")
		return
	}
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func reverse(node *Node) {
	if node == nil {
		fmt.Println("empty")
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	for node.Previous != nil {
		fmt.Printf("%d <- ", node.Value)
		node = node.Previous
	}
	fmt.Println()
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	sum := 0
	for node.Next != nil {
		sum++
		node = node.Next
	}
	return sum
}

func lookup(node *Node, v int) *Node {
	if node == nil {
		return nil
	}
	if v == node.Value {
		return node
	}
	if node.Next == nil {
		return nil
	}
	return lookup(node.Next, v)
}

func main() {
	root := &Node{1, nil, nil}
	fmt.Println(root)
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
