package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("not enought args")
		return
	}
	input := args[1]

	reg := regexp.MustCompile(`([2-3][0-9][0-9]|[4][0][0])`)

	res := reg.FindAllString(input, -1)
	fmt.Println(res)
}
