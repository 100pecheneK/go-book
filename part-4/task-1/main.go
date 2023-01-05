package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: %s.go IPv4\n", filepath.Base(args[0]))
		os.Exit(1)
	}
	ip := args[1]
	ipParts := strings.Split(ip, ".")
	if len(ipParts) != 4 {
		fmt.Println("error IPv4 adress")
		return
	}

	for i, part := range ipParts {
		p, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("err part %d: %s\n", i+1, part)
			continue
		}
		if p < 0 || p > 255 {
			fmt.Printf("err part %d: %s\n", i+1, part)
		}

	}
}
