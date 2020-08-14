package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
	w float64
}

func (s square) area() float64 {
	return s.l * s.w
}

type circle struct {
	r float64
}

func (s circle) area() float64 {
	return math.Pi * math.Pow(s.r, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{10}
	s := square{10, 10}
	info(c)
	info(s)
}
