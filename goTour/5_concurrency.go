// concurrency
package main

import (
	"fmt"
	"sync" // mutex: lock & unlock
	"time"
)

// print out string
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// start new goroutine (thread)
func goroutines() {
	go say("thread: world")
	say("normal: hello")
}

// sum numbers in slice
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

// send/receive data (connect concurrent goroutines)
// data flow in direction of arrow <-
func channels() {
	s := []int{7, 2, 8, -9, 4, 0}

	// unbuffered communication (synchronous communication)
	c := make(chan int) // channels must be created before use
	// distribute work between 2 goroutines (threads)
	go sum(s[:len(s)/2], c)
	go sum(s[len(c)/2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// buffered channels (for asynchronous communication)
	// sends block when buffer is full, receives block when buffer is empty
	ch := make(chan int, 2) // buffer = 2
	ch <- 1
	ch <- 2 // if not have 2 values in chan, then receive causes deadlock in receive
	// ch <- 3 // causes deadlock since channel buffer is already full (send)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// fibonacci with channel
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // send x's value to the channel
		x, y = y, x+y
	}
	close(c) // sender closes channel
}

// sender can close channel to indicate that no more values will be sent
// terminating a channel only necessary if receiver has to know no more values are coming (terminate range loop)
// receiver can test if channel is closed through 2nd argument
func rangeClose() {
	c := make(chan int, 10) // buffer limits number of goroutines launched
	go fibonacci(cap(c), c) // run fibonacci until index 10
	for i := range c {
		fmt.Println(i)
	}
	// TODO: how to run multiple concurrent processes
}

// fibonaci with select
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// TODO: confusion regarding this select situation
func someFunc(c, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

// blocks until one of its cases can run & then executes that case
func selectConcurrent() {
	// TODO: confusion regarding this select situation
	c := make(chan int)
	quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	go someFunc(c, quit)
	fibonacciSelect(c, quit)

	// default selection: run if no other case is ready (avoid blocking)
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// increment counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // lock so that only 1 goroutine can access the map c.v at a time
	c.v[key]++
	c.mu.Unlock()
}

// get current counter value for a given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // defer ensures that the mutex unlocks
	return c.v[key]
}

// avoid conflicts by allowing only one goroutine to access a variable at a time
func syncMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	fmt.Println(c.Value("somekey"))
	time.Sleep(time.Second) // allows computation to complete by waiting
	fmt.Println(c.Value("somekey"))

}

// goroutine = lightweight thread managed by Go runtime
func main() {
	fmt.Println("Concurrency with Goroutines")

	// goroutines()
	// channels()
	// rangeClose()
	// selectConcurrent()
	syncMutex()
}
