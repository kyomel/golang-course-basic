package main

import "fmt"

func main() {
	// x := 83 % 40
	// y := 83 / 40
	// fmt.Println(x)
	// fmt.Println(y)

	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("done.")
}
