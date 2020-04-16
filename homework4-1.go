package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	a float64
}

func (s Square) area() float64 {
	return s.a*s.a
}

func (s Square) perimeter() float64 {
	return 4*s.a
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c Circle) perimeter() float64 {
	return 2*c.radius*math.Pi
}

type Rectangle struct {
	a float64
	b float64
}

func (r Rectangle) perimeter() float64 {
	return 2*(r.a+r.b)
}

func (r Rectangle) area() float64 {
	return r.a*r.b
}


func main() {

	var s Figure = Square{1}
	var c Figure = Circle{1}
	var r Figure = Rectangle{3,4}

	arr := []Figure{s,c,r}
	for _, f:=range arr {
		fmt.Printf("%T\nArea %f\nPerimeter %f\n\n",f,f.area(), f.perimeter())
	}
	
	fmt.Println("Square area is", s.area(), "\nSquare perimeter is", s.perimeter())
	fmt.Println("\nCircle area is", c.area(), "\nCircle perimeter is", c.perimeter())
	fmt.Println("\nRectangle area is", r.area(), "\nRectangle perimeter is", r.perimeter())
}