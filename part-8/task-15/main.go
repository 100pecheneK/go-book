package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	return nil
}

func main() {
	args := os.Args
	err := wordByWord(args[1])
	if err != nil {
		fmt.Println(err)
	}
}
