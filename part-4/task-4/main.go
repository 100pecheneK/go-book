package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Numbers struct {
	P1  int `json:"p1"`
	P2  int `json:"p2"`
	P3  int `json:"p3"`
	P4  int `json:"p4"`
	P5  int `json:"p5"`
	P6  int `json:"p6"`
	P7  int `json:"p7"`
	P8  int `json:"p8"`
	P9  int `json:"p9"`
	P10 int `json:"p10"`
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("usage: %s.go data.json", filepath.Base(args[0]))
		os.Exit(1)
	}
	fileData, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var numbers Numbers
	err = json.Unmarshal(fileData, &numbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	numbers.P1++
	numbers.P2++
	numbers.P3++
	numbers.P4++
	numbers.P5++
	numbers.P6++
	numbers.P7++
	numbers.P8++
	numbers.P9++
	numbers.P10++
	updatedNumbers, err := json.Marshal(numbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(args[1], updatedNumbers, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
