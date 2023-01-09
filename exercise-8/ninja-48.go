package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string   `json:"first"`
	Last   string   `json:"last"`
	Age    int      `json:"age"`
	Saying []string `json:"sayings"`
}

func main() {
	s := `[{"First":"James", "Last": "Bond", "Age": 32, "sayings":["hello", "hi"]}]`

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
