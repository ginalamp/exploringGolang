// more types: pointers, structs, slices, and maps

package main

import (
	"fmt"
)

// holds memory address value. Default value = nil. No pointer arithmetic.
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

func main() {
	fmt.Println("More types: pointers, structs, slices, and maps")

	pointers()

}
