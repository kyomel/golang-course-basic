package main

import "fmt"

func main() {
	if true {
		fmt.Println("01")
	}

	if false {
		fmt.Println("02")
	}

	if !true {
		fmt.Println("03")
	}

	if !false {
		fmt.Println("04")
	}

	if 2 == 2 {
		fmt.Println("05")
	}

	if 2 != 2 {
		fmt.Println("06")
	}

	if !(2 == 2) {
		fmt.Println("07")
	}

	if !(2 != 2) {
		fmt.Println("08")
	}

	if x := 42; x == 42 {
		fmt.Println("001")
	}
	fmt.Println("here's a statement")
	fmt.Println("something else")
}
