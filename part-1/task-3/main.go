package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// считать аргументы
// пока не встретилось в аргументах END считаь в цикле сумму целых чисел

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enought arguments")
		return
	}
	i := 1
	var ints []int
	for {
		if i == len(arguments) {
			fmt.Println("No END")
			break
		}
		if arguments[i] == "END" {
			break
		}
		argument, err := strconv.Atoi(arguments[i])
		if err != nil {
			i++
			continue
		}
		ints = append(ints, argument)
		i++
	}

	fmt.Println(Join(ints))

}

func Join(arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ", "), "[]")
}
