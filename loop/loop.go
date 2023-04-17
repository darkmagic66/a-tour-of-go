package main

import "fmt"

func main() {
	var arr = [...]string{"rust", "haskell", "go"}

	// for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// while loop
	var n int = 0
	for n < 5 {
		fmt.Println(n)
		n += 1
	}

	// for each
	for i, val := range arr {
		fmt.Println(i, val)
	}
	// best practice
	for _, val := range arr {
		fmt.Println(val)
	}
}
