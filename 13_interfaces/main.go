package main

import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	area() float64
}

// define some structs
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func get_area(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{radius: 30}
	rectangle := Rectangle{width: 20, height: 10}

	fmt.Printf("Circle Area is: %f\n", get_area(circle))
	fmt.Printf("rectangle Area is: %f\n", get_area(rectangle))
}