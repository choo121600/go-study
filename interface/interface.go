package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	// perimeter() float64
}

type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Triangle struct {
	width, height float64
}

func (t Triangle) area() float64 {
	return t.width * t.height / 2
}
func main() {
	r := Rect{10.0, 20.0}
	c := Circle{10}
	t := Triangle{10.0, 20.0}
	showArea(r, c, t)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}
