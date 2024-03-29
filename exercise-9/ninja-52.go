package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
	}()

	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
