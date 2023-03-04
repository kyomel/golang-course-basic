package main

import (
	"os"
)

func main() {
	// defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println("err happened", err)
		// log.Fatalln(err)
		panic(err)
	}
}

// func foo() {
// 	fmt.Println("When os.Exit() is called, deffered functions don't run")
// }
