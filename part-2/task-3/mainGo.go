package main

import "C"

import "fmt"

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function")
}

//export Myltiply
func Myltiply(a, b int) int {
	return a * b
}

func main() {

}
