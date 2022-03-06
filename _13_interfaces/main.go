package main

import (
	"fmt"
	"math"
)

//Define interface

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, raduis float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, raduis: 8}
	rectangle := Rectangle{h: 9, w: 9}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
