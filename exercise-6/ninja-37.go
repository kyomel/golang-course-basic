package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Halo Vania")
}

func foo() {
	defer func() {
		fmt.Println("Foo DEFER ran")
	}()
	fmt.Println("Foo ran")
}
