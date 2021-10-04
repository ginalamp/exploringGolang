// packages, variables, and functions
package main

import (
	"fmt" // format
	"math"
	"math/cmplx" // complex numbers
	"math/rand"
	"strconv" // convert element to string
	"time"
)

// package constants
const Pi = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// package variables
var c, python, java bool // initialised to false
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// adds to integers and returns result.
func add(x, y int) int {
	return x + y
}

// swaps given strings and returns result as 2 seperate strings.
func swap(x, y string) (string, string) {
	return y, x
}

// return named values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界! Welcome to the playground")
	fmt.Println("The time is:", time.Now())

	// https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator
	rand.Seed(time.Now().UnixNano()) // set random number seed
	fmt.Println("My fave number is ", rand.Intn(10))

	// print format
	fmt.Printf("now you have %g problems\n", math.Sqrt(7))

	// call function
	x := 1
	y := 2
	fmt.Printf("%d + %d = %d\n", x, y, add(x, y))

	// convert int to string
	fmt.Println(swap(strconv.Itoa(x), strconv.Itoa(y)))

	// split a number into 2 numbers (sum of numbers equal orig num)
	fmt.Println(split(17))

	// variables
	var i int // initialised to 0
	var j, k int = 666, 420
	fmt.Println(i, j, k, c, python, java)

	// basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// type inference
	v := 0.867 + 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)

	// constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// number constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// overflows int (an int can store max 64-bit integer)
	// fmt.Println(needInt(Big))
}
