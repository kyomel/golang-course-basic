package main

import "fmt"

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS
type person struct {
	first string
	last  string
}

type foo int

var y foo

const bar int = 42

func main() {
	y = 42
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)
}
