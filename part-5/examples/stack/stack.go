package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	node *Node
	size int
}

func (stack *Stack) Push(v int) {
	if stack.node == nil {
		stack.node = &Node{v, nil}
		return
	}
	stack.node = &Node{v, stack.node}
	stack.size++
}
func (stack *Stack) Pop() (int, bool) {
	if stack.node == nil {
		return 0, false
	}
	value := stack.node.Value
	stack.node = stack.node.Next
	stack.size--
	return value, true
}

func (stack Stack) traverse() {
	if stack.node == nil {
		fmt.Println("empty")
		return
	}
	for stack.node != nil {
		fmt.Printf("%d -> ", stack.node.Value)
		stack.node = stack.node.Next
	}
	fmt.Println()
}
func main() {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.traverse()
	v, b := stack.Pop()
	if b {
		fmt.Println("Pop: ", v)
	}
	stack.traverse()
	stack.Push(5)
	stack.traverse()

}
