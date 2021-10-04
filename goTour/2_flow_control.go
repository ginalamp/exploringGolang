// basic flow control statements: for, if, else, switch, defer

package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

// basic switch understanding
func switch_statements() {
	fmt.Println("Go runs on ")

	// switch cases do not need to be integers
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// do not need break statement (automatically added)
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// switch cases do not need to be constants
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}

	// switch with no conditions is the same as switch true (ideal for long if-then-else chains)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// basic understanding of defer statements
func defer_statements() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed
	// until the surrounding function returns.

	// will only execute after surrounding function (hello and stack_defer) is done
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defers
	stack_defer()
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
func stack_defer() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		// will print in revers order
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	fmt.Println("Flow control statements: for, if, else, switch, defer")
	// for_loops()
	// if_else()
	// switch_statements()
	defer_statements()

}
