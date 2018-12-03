package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Info", s.area())
}

func main() {
	s := square{
		side: 2.4,
	}
	c := circle{
		radius: 3.2,
	}

	info(s)
	info(c)
}
