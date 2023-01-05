package main

import (
	"fmt"
	"time"
)

// Alternate way
// type Weekday int

func main() {
	const (
		sunday time.Weekday = iota
		monday
		tusday
		wensday
		thursday
		friday
		saturday
	)
	fmt.Println(monday,
		tusday,
		wensday,
		thursday,
		friday,
		saturday,
		sunday,
	)
}

// Alternate way
// func (w Weekday) String() string {
// 	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w]
// }
