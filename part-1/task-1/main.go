package main

// 1) Прочитать аргуменыты командной строки
// 2) Получить действительные числа
// 3) Посчитаь сумму
// 4) Вывести результат

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
	result := 0.0
	for i := 1; i < len(arguments); i++ {
		argument, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}
		result += argument

	}

	fmt.Printf("Result: %f\n", result)

}
