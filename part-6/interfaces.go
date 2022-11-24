package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface{
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func(h *human) {
	fmt.Println("I called human")
}
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1 = person{
		first
	}

	p1 := person{
		fir
	}

	p1 := person{
		first: "Dr."
		last: "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}

