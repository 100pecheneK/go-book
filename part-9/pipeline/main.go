package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var DATA = make(map[int]bool)

var signal chan struct{}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func writeRandomTo(min, max int, out chan<- int) {
	for {
		x := random(min, max)
		select {
		case <-signal:
			fmt.Println("\nclose!")
			close(out)
			return
		case out <- x:
			fmt.Println("\ngenerated: ", x)
		}
	}
}

func checkUniq(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, "-")
		_, ok := DATA[x]
		if ok {
			fmt.Printf("\n%d is already exists\n", x)
			signal <- struct{}{}
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println("close 2")
	close(out)
}

func printSum(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	n1, n2 := getN1N2()
	rand.Seed(time.Now().UnixNano())

	A := make(chan int)
	B := make(chan int)
	signal = make(chan struct{})

	go writeRandomTo(n1, n2, A)
	go checkUniq(B, A)
	printSum(B)

	fmt.Println(DATA)
}

func getN1N2() (n1, n2 int) {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer params")
		return
	}
	n1, _ = strconv.Atoi(os.Args[1])
	n2, _ = strconv.Atoi(os.Args[2])
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	return
}
