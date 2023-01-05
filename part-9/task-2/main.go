package main

import "fmt"

func sumRect(start int, s, quit chan int) {
	counter := start
	sum := counter * counter
	for {
		select {
		case s <- sum:
			// fmt.Printf("%d^2(%d) + %d = %d\n", counter+1, (counter+1)*(counter+1), counter, (counter+1)*(counter+1)+counter)
			sum += (counter + 1) * (counter + 1)
			counter++
		case <-quit:
			fmt.Println("quit")
			close(s)
			return
		}
	}
}

func main() {
	s := make(chan int)
	quit := make(chan int)
	start := 1
	go func() {
		for i := start; i <= 3; i++ {
			fmt.Println(<-s)
		}
		quit <- 0
	}()
	sumRect(start, s, quit)

}
