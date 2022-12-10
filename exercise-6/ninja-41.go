package main

import "fmt"

//declaration asignment
var x int = 7
var g func() = func() {
	fmt.Println("g from outside")
}

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g()
	g = f
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}
