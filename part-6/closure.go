package main

import "fmt"

var x int

func main() {
	// if x inside make error
	// fmt.Println(x)
	// x++
	// fmt.Println(x)

	// example error code block with code block
	// {
	// 	y := 42
	// 	fmt.Println(y)
	// }
	// fmt.Println(y)

	// foo()
	// fmt.Println(x)

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// func foo() {
// 	fmt.Println("hello")
// 	x++
// }

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
