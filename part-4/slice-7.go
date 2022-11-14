package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "JackDaniels"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "BaliHai"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
