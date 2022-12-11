package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)

	c := &a
	fmt.Println(c)
	fmt.Println(*c) // * givers you the value stored at an address when you have the address
	fmt.Println(*&a)

	*c = 43
	fmt.Println(a)
}
