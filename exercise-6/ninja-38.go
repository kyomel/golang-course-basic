package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I'm", p.first, p.last, "am I'm", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "Francesca",
		last:  "Vania",
		age:   24,
	}

	p1.speak()
}
