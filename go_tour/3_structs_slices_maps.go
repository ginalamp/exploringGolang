// more types: pointers, structs, arrays, slices, and maps

package main

import (
	"fmt"
)

// struct: collection of fields
type Vertex struct {
	X int
	Y int
}

// struct literal
var (
	v1 = Vertex{1, 2}  // type Vertex
	v2 = Vertex{X: 1}  // {1, 0}
	v3 = Vertex{}      // {0, 0}
	px = &Vertex{1, 2} // type *Vertex
)

// pointers holds memory address value. Default value = nil. No pointer arithmetic.
func pointers() {
	i, j := 42, 2701

	p := &i         // p points to i
	fmt.Println(*p) // dereference p: read i's value
	*p = 21         // set p's (i's) value to 21
	fmt.Println(i)  // see new value of i
	fmt.Println(*p)

	p = &j         // p points to j
	*p = *p / 37   // divide p (j) through 37
	fmt.Println(j) // see new value of j
	fmt.Println(*p)
}

// struct = collection of fields
func structs() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// change struct value
	v.X = 4
	fmt.Println(v)

	// pointers to structs
	p := &v // p point to v
	p.Y = 1e9
	fmt.Println(v)

	// struct literal
	fmt.Println(v1, v2, v3, px)
}

func main() {
	fmt.Println("More types: pointers, structs, slices, and maps")

	// pointers()
	structs()

}
