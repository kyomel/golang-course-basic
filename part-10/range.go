package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // close tidak digunakan maka akan menyebabkan deadlock
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}
