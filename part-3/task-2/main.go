package main

import "fmt"

func main() {
	arrayI := []int{1, 2, 3, 4}
	arrayS := []string{"one", "two", "three", "four"}
	i := convertArrayToHashTable(arrayI)
	s := convertArrayToHashTable(arrayS)
	fmt.Println(i)
	fmt.Println(s)
}

func convertArrayToHashTable[T any](arr []T) map[int]T {
	m := make(map[int]T)
	for i, e := range arr {
		m[i] = e
	}
	return m
}
