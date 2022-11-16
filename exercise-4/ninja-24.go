package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	x = append(x, 51)
	fmt.Println(x)
	x = append(x, 52, 53, 54)
	fmt.Println(x)
	y := []int{55, 56, 57, 58, 59}
	x = append(x, y...)
	fmt.Println(x)
}
