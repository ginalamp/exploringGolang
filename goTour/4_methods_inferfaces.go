// methods and interfaces

package main

import (
	"fmt"
	"math"
)

// non-struct type declaration (can only have methods of types within the same package)
type MyFloat float64

// methods have a special receiver argument
// func <receiver> <functionName>() <returnType> {}
func methods() {
	basics()
	pointerReceivers()
}

// vertex with (x,y) coordinates
type Vertex struct {
	X, Y float64
}

// calling methods
func basics() {
	// create vertex and compute value
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// absolute value of non-struct type
	f := MyFloat(-math.Sqrt2) // convert float to MyFloat
	fmt.Println(f.Abs())
}

// get Vertex square root of (X^2 + Y^2)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// get the absolute value of f
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// better performance (avoid copying values)
// methods can receive pointers (not for pointers to pointers)
// can modify the value to which the receiver points
// can send value or pointer to method (automatically converts)
func pointerReceivers() {
	// send Vertex value
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// send pointer to Vertex
	p := &Vertex{4, 3}
	p.Scale(3)
	fmt.Println(p.Abs())
}

// receives pointer and modifies values to multiply with f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println("Running methods & interfaces")
	methods()
}
