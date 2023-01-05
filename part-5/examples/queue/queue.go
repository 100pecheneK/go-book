package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type Queue struct {
	node *Node
	size int
}

func (queue *Queue) Push(v int) {
	if queue.node == nil {
		queue.node = &Node{v, nil}
		return
	}
	queue.node = &Node{v, queue.node}
	queue.size++
}

func (queue Queue) Pop() (int, bool) {
	if queue.size == 0 {
		return 0, false
	}
	if queue.size == 1 {
		queue.size--
		return queue.node.Value, true
	}
	var newLastNode *Node
	for queue.node.Next != nil {
		newLastNode = queue.node
		queue.node = queue.node.Next
	}

	oldLastNodeValue := newLastNode.Next.Value
	newLastNode.Next = nil

	queue.size--
	return oldLastNodeValue, true
}

func (queue Queue) traverse() {
	if queue.node == nil {
		fmt.Println("empty")
		return
	}
	for queue.node != nil {
		fmt.Printf("%d -> ", queue.node.Value)
		queue.node = queue.node.Next
	}
	fmt.Println()
}

func main() {
	queue := new(Queue)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.traverse()
	fmt.Println("Size: ", queue.size)
	poped, isPoped := queue.Pop()
	if isPoped {
		fmt.Println("Pop: ", poped)
	}
	queue.traverse()
	queue.Push(4)
	queue.traverse()

}
