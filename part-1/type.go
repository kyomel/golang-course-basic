package main

import "fmt"

var y = 42

// declare the variable with the identfier "z"
// is of type string
// and I'm assign the value "Shaken, not stirred"

var z = "Shaken, not stirred"
var a string = `James said, 
"Shaken, not stirred"`

// this is a static programming language
// a variable is declared to hold a value of a certain type
// not a dynamic programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Println("%T\n", z)
	fmt.Println(a)
	fmt.Println("%T\n", a)
	// z = 43
	// fmt.Println(z)
	// fmt.Println("%T\n", z)
}
