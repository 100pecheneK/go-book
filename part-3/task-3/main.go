package main

import "fmt"

// y << x - means "y" myltiply 2 "x" times

func main() {
	const (
		c4 = 4 << (2 * iota)
		c16
		c64
		c256
		c1024
	)

	fmt.Println(c4, c16, c64, c256, c1024)
}
