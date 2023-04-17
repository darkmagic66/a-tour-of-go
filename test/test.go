package main

import (
	"fmt"
)

func main() {
	value, err := fmt.Println("Hello Go")
	fmt.Println(value, err) // --> 9 byte is use , nil (no err)
}
