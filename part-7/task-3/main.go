package main

import "fmt"

type Cord struct {
	X, Y int
}

type Distance interface {
	printDistance()
}
type Line struct {
	A, B Cord
}

func (l Line) printDistance() {
	fmt.Printf("%d\n", l.A.X-l.B.X+l.A.Y-l.B.Y)
}

func main() {
	printCords(Line{Cord{1, 2}, Cord{2, 1}})
}

func printCords(line Line) {
	line.printDistance()
}
