package main

import (
	"fmt"
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
}

func main() {
	fmt.Println("Flow control statements: for, if, else, switch, defer")
	for_loops()

}
