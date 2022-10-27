package main

import "fmt"

var x bool

func main() {
	a := 7
	b := 42
	fmt.Println(x)
	x = true
	fmt.Println(x)

	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
}
