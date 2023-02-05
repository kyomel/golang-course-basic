package main

import "fmt"

func main() {
	// c := make(chan int, 1)

	// c <- 42
	// fmt.Println(<-c)

	c := make(chan int, 2)
	// akan berhasil jika buffer sesuai dengan value channel
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
