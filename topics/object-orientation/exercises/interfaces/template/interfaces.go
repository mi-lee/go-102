// Define an interface which defines a method area().  Create types for square,
// rectangle and circle, and ensure they satisfy your interface.  Create a
// function that accepts a value of your interface type and outputs the area,
// and call this function for different shapes.
//
// Template available at: http://play.golang.org/p/rL5tT2VTJH
package main

// Add your imports here.
import (
	"fmt"
	"math"
)

// Define an interface with a method area().  Make sure you use a meaningful
// name and a sensible return type.

type areaGetter interface {
	area() float64
}

// Create square, rectangle and circle types, and ensure they satisfy your
// interface (you'll need to use the `Pi` constant from the `math` package for
// your circle).

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

// Write a function that accepts a value of your interface and outputs the
// area.

func main() {
	// Create a slice of your interface type, and populate it with a number of
	// different shapes.
	stats := make([]areaGetter{circle{2}, square{4}})

	for _, s := range stats {
		fmt.Printf("%T:", s.area())
	}

	// Loop through your shapes and use your function to output the area of
	// each one.
}
