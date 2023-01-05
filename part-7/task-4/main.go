package main

import "fmt"

type X struct {
	value int
}

func (x *X) increment() {
	x.value++
}
func (x *X) print() {
	fmt.Println(x.value)
}
func main() {
	x := X{0}
	x.increment()
	x.increment()
	x.increment()
	x.increment()
	x.print()
}
