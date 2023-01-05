package main

import "fmt"

func main() {
	willClose := make(chan int, 10)

	willClose <- 2
	fmt.Println(<-willClose)

	close(willClose)

	// read := <-willClose
	// fmt.Println(read)
}
