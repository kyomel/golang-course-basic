package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// JSON Unmarshall jika struct yg datanya huruf kecil di dalam makan pada saat unmarshall datanya gak akan muncul

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Francesca",
		Last:  "Vania",
		Age:   24,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
