package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "MoneyPenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BondBondBond", x)
	} else {
		fmt.Println("neither")
	}
}
