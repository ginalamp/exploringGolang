// more types: pointers, structs, arrays, slices, and maps

package main

import (
	"fmt"
	"strings"
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

// Type [n]T is an array of n values of type T.
// Arrays have a fixed size
func arrays() {
	var a [2]string // a = array of 2 strings. Can't resize.
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 666, 721} // int arr of size 6
	fmt.Println(primes)
}

// Slices are dynamically sized, flexible view into an array
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array
	var s []int = primes[1:4]            // slice [3,5,7]
	fmt.Println(s)

	// slice is like a pointer to a section in an array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // [John, Paul]
	b := names[1:3] // [Paul, George]

	// changing an element in a slice changes the corresponding array element
	b[0] = "XXX"
	fmt.Println(a, b) // [John, XXX], [XXX, George]

	// Slice literals: an array without the length
	sliceLiterals()

	// Slice default equivalent expressions
	slice_eq := primes[0:6]
	fmt.Println(slice_eq)
	slice_eq = primes[:6] // lower bound = 0
	fmt.Println(slice_eq)
	slice_eq = primes[0:] // upper bound = slice length
	fmt.Println(slice_eq)
	slice_eq = primes[:]
	fmt.Println(slice_eq)

	// length and capacity
	sliceLenCap()

	// dynamic slices: make
	makeDynamicSlice()

	// slices of slices
	sliceOfSlices()

	// add to slice
	sliceAppend()

	// range returns (<elementIndex>, <elementValue>)
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}
	// skip range return value with _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	for index := range pow {
		pow[index] = 1 << uint(index) // 2**<index>
	}

}

// slice literals standard and struct types
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// slice of temporary structs
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// f()
}

// length = num elements in slice
// capacity = num elements in underlying array (starting from 1st val in slice)
func sliceLenCap() {
	// can change slice length through reslicing if have enough capacity
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	a := s[:0]
	printSlice(a)

	// Extend its length.
	a = s[:4]
	printSlice(a)

	// Drop its first two values.
	a = s[2:]
	printSlice(a)

	fmt.Println(s)
}

// dynamically-sized arrays
// make() allocates a zeroed array & returns a slice that refers to that array
func makeDynamicSlice() {
	a := make([]int, 5) // len(a) = 5 (default capacity = len if not specified)
	printSlice2("a", a)

	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	printSlice2("b", b)

	c := b[:2] // len(c) = 2, cap(c) = cap(b) = 5
	printSlice2("c", c)

	// TODO: why is cap(d) = 3 and not 4 (5-1)?
	d := c[2:5] // len(d) = 3, cap(d) = 3
	printSlice2("d", d)

}

// slice can contain any type (including other slices)
func sliceOfSlices() {
	// create tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// print board
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// add new elements to slice
// If backing array of slice is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
func sliceAppend() {
	// append(<slice>, <valuesToAppendToSlice>)

	var s []int
	printSlice(s)

	// append individual elements
	s = append(s, 0)
	s = append(s, 1)
	printSlice(s)

	// append multiple elements
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func main() {
	fmt.Println("More types: pointers, structs, slices, and maps")

	// pointers()
	// structs()
	// arrays()
	slices()

}

// print slice with it's length and capacity values
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// print slice with it's length and capacity values with var name
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
