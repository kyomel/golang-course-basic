package main

import "fmt"

// using iota
// const (
// 	_  = iota
// 	kb = 1024
// 	mb = 1024 * kb
// 	gb = 1024 * mb
// )

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	// x := 4
	// fmt.Printf("%d\t\t%b\n", x, x)

	// y := x << 1
	// fmt.Printf("%d\t\t%b\n", y, y)

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	// fmt.Println(kb)
	// fmt.Println(mb)
	// fmt.Println(gb)
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
