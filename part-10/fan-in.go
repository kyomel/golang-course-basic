package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	even := make(chan int)
// 	odd := make(chan int)
// 	fanin := make(chan int)

// 	go send(even, odd)

// 	go receive(even, odd, fanin)

// 	for v := range fanin {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("about to exit")

// }

// func send(even, odd chan<- int) {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			even <- i
// 		} else {
// 			odd <- i
// 		}
// 	}
// 	close(even)
// 	close(odd)
// }

// func receive(even, odd <-chan int, fanin chan<- int) {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		for v := range even {
// 			fanin <- v
// 		}

// 		wg.Done()
// 	}()

// 	go func() {
// 		for v := range odd {
// 			fanin <- v
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	close(fanin)
// }

// rob pike code fan in
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm Leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
