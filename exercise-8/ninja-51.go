package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Last:  "Sister",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Last:  "Hmmm",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
	}

	fmt.Println("---------")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
	}

	fmt.Println("---------")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
	}
}
