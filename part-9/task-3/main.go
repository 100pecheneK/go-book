package main

import (
	"fmt"
	"sync"
)

func function(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go function(&waitGroup)
	waitGroup.Add(1)
	go func() {
		waitGroup.Done()
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	// time.Sleep(1 * time.Second)
	waitGroup.Wait()
	fmt.Println("end")
}
