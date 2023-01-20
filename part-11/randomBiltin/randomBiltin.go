package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	type point3D struct {
		X, Y, Z int8
		S       float32
	}
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	values := reflect.TypeOf(point3D{})
	x, _ := quick.Value(values, ran)
	fmt.Println(x)
}
