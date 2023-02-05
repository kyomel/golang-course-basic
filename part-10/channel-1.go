package main

import "fmt"

// channel doesn't run deadlock
func main() {
	// deadlock
	// c := make(chan int)

	// c <- 42

	// fmt.Println(<-c)

	// unsuccessfull buffer
	c := make(chan int, 1)

	c <- 42
	c <- 43
	fmt.Println(<-c)
}
