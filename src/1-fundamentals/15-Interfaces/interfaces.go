package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func calculateArea(s shape) {
	fmt.Printf("The area of the shape is: %0.2f\n", s.area())
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func main() {

	rect := rectangle{10, 20}
	tri := triangle{60, 40}
	circ := circle{10}

	calculateArea(rect)
	calculateArea(tri)
	calculateArea(circ)

	return
}
