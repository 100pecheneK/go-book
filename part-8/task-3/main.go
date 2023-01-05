package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt, syscall.SIGTERM:
				handle(sig)
				os.Exit(0)
			case syscall.SIGINFO:
				handle(sig)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
