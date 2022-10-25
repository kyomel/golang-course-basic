package main

import "fmt"

var a int

// CREATE own type
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T", a)
	fmt.Println(b)
	fmt.Printf("%T", b)

	// conversion, not casting
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
