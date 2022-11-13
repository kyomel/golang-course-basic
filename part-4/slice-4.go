package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 13}
	fmt.Println(x)
	x = append(x, 18, 20, 23)
	fmt.Println(x)

	y := []int{234, 456, 789}
	x = append(x, y...) // variadic parameter
	fmt.Println(x)
}
