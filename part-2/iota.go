package main

import "fmt"

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// other way declaration
const (
	a = iota
	b
	c
)

// iota mean can increment at const variable if outside iota must be reset at 0
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
