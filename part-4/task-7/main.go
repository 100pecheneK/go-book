package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type application struct {
	logger *log.Logger
}

const LOGFILE = "./error.log"

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func (app *application) ADD(k string, n myElement) bool {
	if k == "" {
		app.logger.Printf("ADD - unknown key string")
		return false
	}
	if app.LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	app.logger.Printf("ADD - key \"%s\" allready exists", k)
	return false
}

func (app *application) DELETE(k string) bool {
	if app.LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	app.logger.Printf("DELETE - unknow key: %s\n", k)
	return false
}

func (app *application) LOOKUP(k string) *myElement {
	n, ok := DATA[k]
	if !ok {
		app.logger.Printf("LOOKUP - unknow key: %s\n", k)
		return nil
	}
	return &n
}

func (app *application) CHANGE(k string, n myElement) bool {
	_, ok := DATA[k]
	if !ok {
		app.logger.Printf("CHANGE - unknow key: %s\n", k)
		return false
	}

	DATA[k] = n
	return true
}
func (app *application) PRINT() {
	for k, d := range DATA {
		fmt.Printf("key: %s value %v\n", k, d)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	app := &application{
		logger: logger,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		tokens := strings.Fields(text)
		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "", "", "", "")
		case 2:
			tokens = append(tokens, "", "", "")
		case 3:
			tokens = append(tokens, "", "")
		case 4:
			tokens = append(tokens, "")
		}
		switch tokens[0] {
		case "PRINT":
			app.PRINT()
		case "STOP":
			return
		case "DELETE":
			if !app.DELETE(tokens[1]) {
				fmt.Println("delete failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !app.ADD(tokens[1], n) {
				fmt.Println("add failed!")
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !app.CHANGE(tokens[1], n) {
				fmt.Println("change failed!")
			}
		default:
			app.logger.Printf("unknow command: %s\n", tokens[0])
			fmt.Println("unknow command!")
		}
	}
}
