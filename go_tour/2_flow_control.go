package main

import (
	"fmt"
	"math"
)

// basic for loop understanding
func for_loops() {
	// basic for loop (go only has for loops)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for can be like a while loop => optional init and post statements
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop if exit condition not specified
	// for {
	// }

	// for loops and functions
	Sqrt(1)
}

// Newton's method: compute sqrt using loop through guesses
func Sqrt(x float64) float64 {
	z := x / 2
	z_prev := -z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		// break if value not changing
		if z == z_prev {
			fmt.Println("Exiting loop")
			break
		}
		z_prev = z
		fmt.Println(z)
	}
	fmt.Printf("The sqrt of %g is ~%g\n", x, z)
	return z
}

// basic if/else understanding
func if_else() {
	fmt.Println(sqrt(2), sqrt(-4))
	// Both calls to pow return their results before the call to fmt.Println begins
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// get sqrt of float as string
func sqrt(x float64) string {
	// irrational number
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// Sprint converts number to string
	return fmt.Sprint(math.Sqrt(x))
}

// get biggest number between x^n or lim and return it
func pow(x, n, lim float64) float64 {
	// if statement can have short statement to execute before start of condition
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// will print out before the numbers, since pow calls complete before it is printed out in if_else()
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println("Flow control statements: for, if, else, switch, defer")
	for_loops()
	// if_else()

}
