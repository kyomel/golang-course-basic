package main

import "fmt"

func main() {
	// s1 := foo()
	// fmt.Println(s1)

	// x := func() int {
	// 	return 451
	// }()

	x := bar()
	fmt.Printf("%T\n", x)

	// i := x() // clean
	fmt.Println(x())
}

func foo() string {
	// clean code
	// s := "Hello World"
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
