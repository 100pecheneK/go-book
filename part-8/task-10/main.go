package main

import (
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
	return nil
}

func main() {
	filename := ""
	args := os.Args
	for i := 1; i < len(args); i++ {
		filename = args[i]
		fmt.Println(filename)
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
