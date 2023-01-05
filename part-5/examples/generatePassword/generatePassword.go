package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix())
	LENGTH := 8

	startChar := "!"

	i := 0
	for i != LENGTH {
		char := string(startChar[0] + byte(random(0, 94)))
		fmt.Print(char)
		i++
	}

	fmt.Println()
}
