package main

type Shape3D interface {
	calculateArea() int
}

type Cube struct {
	X, Y, Z int
}

type Sphere struct {
	D, R int
}

func (s Sphere) calculateArea() int {
	return s.D * s.R
}

func (c Cube) calculateArea() int {
	return c.X * c.Y * c.Z
}
