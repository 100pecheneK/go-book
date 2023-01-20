package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	domain := os.Args[1]

	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
