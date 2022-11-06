package main

import "fmt"

func main() {
	// for single condition
	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// fmt.Println("done.")

	//  for condition if
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("done.")
}
