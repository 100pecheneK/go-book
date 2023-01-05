package main

import (
	"fmt"
	"t/internal/handlers"
	"t/internal/utils"
)

func main() {
	a := handlers.Handle(2)
	fmt.Println(a)
	fmt.Println(utils.Sum(a, 2))
}
