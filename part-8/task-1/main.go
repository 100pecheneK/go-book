package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filepath, search, replace, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = searchAndReplaceInFile(filepath, search, replace)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func getArgs() (filepath, search, replace string, err error) {
	args := os.Args
	if len(args) < 4 {
		return "", "", "", errors.New("not enought arguments\nneed args: filepath search replace")
	}
	filepath, search, replace, err = args[1], args[2], args[3], nil
	return
}

func searchAndReplace(string, search, replace string) string {
	s := strings.ReplaceAll(string, search, replace)
	return s
}
func writeString(w io.Writer, string string) {
	io.WriteString(os.Stdout, string)
	io.WriteString(os.Stdout, "\n")
}
func writeStringToStdout(string string) {
	writeString(os.Stdout, string)
}

func searchAndReplaceInFile(filepath, search, replace string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		replaced := searchAndReplace(scanner.Text(), search, replace)
		writeStringToStdout(replaced)
	}
	return nil
}
