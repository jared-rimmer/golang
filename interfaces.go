package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g shape) {
	fmt.Println(g)
	fmt.Println(g.area())
}

// func main() {
// 	r := rectangle{width: 3, height: 4}
// 	c := circle{radius: 5}

// 	measure(r)
// 	measure(c)
// }
