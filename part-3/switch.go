package main

import "fmt"

func main() {
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2 == 4):
	// 	fmt.Println("this should not print 2")
	// case (3 == 3):
	// 	fmt.Println("prints")
	// 	fallthrough // digunakan utk mencari ke case selanjutnya yg bernilai true atau false ditampilkan
	// case (4 == 4):
	// 	fmt.Println("also true, does it print?")
	// 	fallthrough
	// case (7 == 9):
	// 	fmt.Println("not true 1")
	// 	fallthrough
	// case (11 == 14):
	// 	fmt.Println("not true 2")
	// 	fallthrough
	// case (15 == 15):
	// 	fmt.Println("true 15")
	// 	fallthrough
	// default:
	// 	fmt.Println("this is default")
	// }

	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2 == 4):
	// 	fmt.Println("this should not print 2")
	// default:
	// 	fmt.Println("this is default")
	// }

	n := "Bond"
	switch n {
	case "Moneypenny", "Q", "T":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	default:
		fmt.Println("this is default")
	}

}
