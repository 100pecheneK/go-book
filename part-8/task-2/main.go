package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

const SIZE = 15
const FILEPATH = "./data.gob"

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

func load[T any](filepath string, data *T) error {
	fmt.Println("Loading", filepath)
	loadFrom, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Empty!")
		return err
	}
	defer loadFrom.Close()

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&data)
	return nil
}

func save[T any](filepath string, data *T) error {
	fmt.Println("Saving", filepath)
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Cannot create", filepath)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(*data)
	if err != nil {
		fmt.Println("Cannot save to", filepath)
		return err
	}
	return nil
}

type User struct {
	Name string
}

func main() {
	table := make(map[int]*Node, 4)
	hashTable := &HashTable{table, 4}
	fmt.Println("Number of spaces:", hashTable.Size)
	for i := 0; i < 39; i++ {
		insert(hashTable, i)
	}
	traverse(hashTable)
	save(FILEPATH, hashTable)
	for i := 0; i < 9; i++ {
		insert(hashTable, i)
	}
	fmt.Println("Editing hashTable without saving")
	traverse(hashTable)

	ltable := make(map[int]*Node, 4)
	lhashTable := &HashTable{ltable, 4}
	traverse(lhashTable)
	load(FILEPATH, lhashTable)
	traverse(lhashTable)

	fmt.Println("\n----Users slice----")

	users := []User{
		{
			Name: "1",
		},
		{
			Name: "2",
		},
	}
	fmt.Println(users)
	save(FILEPATH, &users)

	fmt.Println("Editing users without saving")
	users = append(users, User{Name: "3"})
	fmt.Println(users)

	lusers := []User{}
	load(FILEPATH, &lusers)
	fmt.Println(lusers)
}
