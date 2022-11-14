package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "Miss Moneypenny")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["James"]; ok {
		fmt.Println("value:", v)
		delete(m, "James")
	}

	fmt.Println(m)
}
