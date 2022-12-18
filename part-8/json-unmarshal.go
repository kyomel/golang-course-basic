package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `
		[{"First":"James","Last":"Bond","Age":32},{"First":"Francesca","Last":"Vania","Age":24}]
	`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []person{}
	// var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)

	}
}
