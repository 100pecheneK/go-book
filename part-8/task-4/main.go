package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := flag.String("f", "", "file")
	spaces := flag.Int("s", 1, "spaces count")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = replaceTab(line, *spaces)
		fmt.Println(line)
	}
}

func replaceTab(line string, spaces int) string {
	line = strings.ReplaceAll(line, "\t", strings.Repeat(" ", spaces))
	return line
}
