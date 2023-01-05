package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("not enought args")
		return
	}
	input := args[1]

	reg := regexp.MustCompile(`.*(\d\d \d\d \d\d\d\d \d\d:\d\d).*`)
	if !reg.MatchString(input) {
		fmt.Println("no date time in input arg")
		return
	}
	dateTime := reg.FindStringSubmatch(input)
	fmt.Println(dateTime)
	t, err := time.Parse("02 01 2006 15:04", dateTime[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hours: ", t.Hour(), "Minutes: ", t.Minute())

}
