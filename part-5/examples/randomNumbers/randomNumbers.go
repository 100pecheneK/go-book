package main

import (
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(3)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", random(0, 5))
	}
	fmt.Println()
}
