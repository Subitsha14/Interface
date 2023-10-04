package main

import (
	"fmt"
	"math"
)

type calculate interface {
	area() float64
	perim() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func value(a interface{}) {
	val, ok := a.(string) //type assertion
	fmt.Println(val, ok)
}

func main() {
	fmt.Println("Interfaces")
	var r calculate //r and c is of type calculate interface
	var c calculate
	r = rect{height: 10, width: 12}
	c = circle{radius: 20}

	fmt.Println(r.area())
	fmt.Println(r.perim())
	fmt.Println(c.area())
	fmt.Println(c.perim())

	var val interface{} = "Hello"
	value(val)
}
