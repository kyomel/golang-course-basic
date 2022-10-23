package main

import "fmt"

// declare and assign variable "y" with valye 43
// declare & assign = intialization
var y = 43

// declare there is a variable with the identifier "z" and that the variable with the identifier "z" is of type int
// assignt the zero value of type int into "z"
// false for boolean, 0 for integers, 0.0 for floats, "" for strings
// and nil for pointers, functions, interfaces, slices, channels, and maps
var z int

func main() {
	// short declaration operator
	// Declare a variable and Assign a value(of a certain type)
	// suggest : using var at outside function and using short declaration inside function
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
