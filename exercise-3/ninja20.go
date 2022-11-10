package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "play game":
		fmt.Println("go to room!")
	case "basketball":
		fmt.Println("go to court!")
	case "soccer":
		fmt.Println("go to field!")
	}
}
