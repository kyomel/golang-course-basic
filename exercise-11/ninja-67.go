package main

import "fmt"

type hotdog int

// part 1
// type customErr struct {
// 	info string
// }

// func (ce customErr) Error() string {
// 	return fmt.Sprintf("here is the error: %v", ce.info)
// }

func main() {
	// c1 := customErr{
	// 	info: "need more coffe",
	// }

	// foo(c1)

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

// func foo(e error) {
// 	fmt.Println("foo ran -", e, "\n")
// 	fmt.Println("foo ran -", e, "\n", e.(customErr).info) // static programming
// }
