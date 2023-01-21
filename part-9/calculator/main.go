package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Expr struct {
	A int
	B int
	T string
}

func main() {
	input := os.Args[1]
	result := calc(input)
	fmt.Println(result)
}

func calc(input string) string {
	var out string
	ch := make(chan string)
	exprs := make(map[int]Expr)

	parse(input, exprs)

	for _, e := range exprs {
		switch e.T {
		case "+":
			go plus(e.A, e.B, ch)
		case "-":
			go minus(e.A, e.B, ch)
		case "*":
			go mult(e.A, e.B, ch)
		case "/":
			go div(e.A, e.B, ch)
		}
	}

	for i := 0; i < len(exprs); i++ {
		out += fmt.Sprintln(<-ch)
	}
	return out
}

func plus(a, b int, ch chan string) {
	ch <- fmt.Sprintf("%d + %d = %d", a, b, a+b)
}
func minus(a, b int, ch chan string) {
	ch <- fmt.Sprintf("%d - %d = %d", a, b, a-b)
}
func mult(a, b int, ch chan string) {
	ch <- fmt.Sprintf("%d * %d = %d", a, b, a*b)
}
func div(a, b int, ch chan string) {
	ch <- fmt.Sprintf("%d / %d = %d", a, b, a/b)
}

func parse(input string, exprs map[int]Expr) {

	in := strings.Split(input, ",")

	for i, e := range in {
		reg := `(\d*)([\+\-\*\/])(\d*)`
		match := regexp.MustCompile(reg)
		parts := match.FindAllStringSubmatch(e, -1)[0]
		a, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
		}
		b, _ := strconv.Atoi(parts[3])
		expr := Expr{
			A: a,
			B: b,
			T: parts[2],
		}
		exprs[i] = expr
	}
}
