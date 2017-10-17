package main

import "fmt"
import "math"

type square struct {
	side int
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return float64(s.side * s.side)
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%f\n", s.area())
}

func main() {
	// var s shape
	s := square{side: 12}
	// s = nil
	s.side = 12
	c := circle{radius: 12}
	info(s)
	info(c)

}
