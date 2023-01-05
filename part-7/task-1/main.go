package main

import "fmt"

type I interface {
	toPrint() string
}

type S struct {
	name string
	age  int
}

type S2 struct {
	name   string
	age    int
	weight int
}

func (s S) toPrint() string {
	return fmt.Sprintf("Name: %s, Age: %d\n", s.name, s.age)
}
func (s S2) toPrint() string {
	return fmt.Sprintf("Name: %s, Age: %d, Weight: %d\n", s.name, s.age, s.weight)
}

func main() {
	s := S{
		"Misha",
		21,
	}
	s2 := S2{
		"Misha",
		21,
		80,
	}

	print(s)
	print(s2)
}

func print(i I) {
	fmt.Print(i.toPrint())
}
