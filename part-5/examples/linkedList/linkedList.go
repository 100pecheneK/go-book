package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (node *Node) addNode(v int) *Node {
	if v == node.Value {
		fmt.Println("Node already exist:", v)
		return node
	}
	if node.Next == nil {
		node.Next = &Node{v, nil}
		return node.Next
	}
	return node.Next.addNode(v)
}

func traverse(node *Node) {
	if node == nil {
		fmt.Println("Empty list!")
		return
	}

	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func lookupNode(node *Node, v int) *Node {
	if node == nil {
		return nil
	}
	for node.Value != v {
		node = node.Next
	}
	return node
}

func (node *Node) size() int {
	sum := 1
	for node.Next != nil {
		sum++
		node = node.Next
	}
	return sum
}

func main() {
	rootNode := &Node{1, nil}
	fmt.Println("size:", rootNode.size())
	traverse(rootNode)
	rootNode.addNode(2).addNode(3)
	rootNode.addNode(4)
	traverse(rootNode)
	node := lookupNode(rootNode, 4)
	fmt.Println(node)
	node.addNode(5)
	traverse(rootNode)
	node = lookupNode(rootNode, 1)
	fmt.Println(node)
	fmt.Println("size:", rootNode.size())
}
