package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	value int
	data  any
}
type Heaps []*Heap

func (n *Heaps) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	// if pop adress
	// old[len(old)-1] = nil
	new := old[0 : len(old)-1]
	*n = new
	return x
}
func (n *Heaps) Push(x interface{}) {
	*n = append(*n, x.(*Heap))
}

func (n Heaps) Len() int {
	return len(n)
}

func (n Heaps) Less(a, b int) bool {
	return n[a].value < n[b].value
}
func (n Heaps) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &Heaps{
		&Heap{1, 1.2},
		&Heap{2, [2]int{1, 2}},
		&Heap{0, 3.1},
		&Heap{4, "-100.1"},
	}
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)
	heap.Push(myHeap, &Heap{5, "a"})
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)
}
