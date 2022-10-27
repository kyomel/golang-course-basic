package main

import (
	"fmt"
	"runtime"
)

var x int8 = -128

// var y float64

func main() {
	// x = 42
	// y = 42.34534

	fmt.Println(x)
	// fmt.Println(y)
	fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
