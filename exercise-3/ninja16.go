package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (modulus) is %v\n", i, i%4)
	}
}
