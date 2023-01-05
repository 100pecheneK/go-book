package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func insert(hashTable *HashTable, value int) int {
	index := hashFunction(value, hashTable.Size)
	node := Node{value, hashTable.Table[index]}
	hashTable.Table[index] = &node
	return index
}

func traverse(hashTable *HashTable) {
	for k := range hashTable.Table {
		if hashTable.Table[k] != nil {
			node := hashTable.Table[k]
			for node != nil {
				fmt.Printf("%d -> ", node.Value)
				node = node.Next
			}
			fmt.Println()
		}
	}
}

func main() {
	table := make(map[int]*Node, 4)
	hashTable := &HashTable{table, 4}
	fmt.Println("Number of spaces:", hashTable.Size)
	for i := 0; i < 39; i++ {
		insert(hashTable, i)
	}
	traverse(hashTable)
}
