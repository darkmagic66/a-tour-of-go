package main

import (
	"encoding/json"
	"fmt"
)

// struct tag
type Course struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func main() {
	fmt.Println("test")

	c := Course{
		Name:       "basic go",
		Level:      "normal",
		Views:      7239,
		Instructor: "อนุชิโตะ",
		FullPrice:  9999,
	}

	test, _ := json.Marshal(c)
	fmt.Println(test)
	fmt.Printf("%s", test)
}
