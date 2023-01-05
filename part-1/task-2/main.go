package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enought arguments")
		return
	}
	sum := 0.0
	count := 0.0
	for i := 1; i < len(arguments); i++ {
		argument, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		sum += argument
		count++
	}
	fmt.Println("Average: ", sum/count)
}
