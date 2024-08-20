package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

type shape interface {
	getArea() float64
}

func printArea(sh shape) {
	fmt.Printf("Area: %v\n", sh.getArea())
}

func main() {
	tr := triangle{height: 10, base: 5}
	sq := square{sideLength: 10}

	printArea(tr)
	printArea(sq)
}
