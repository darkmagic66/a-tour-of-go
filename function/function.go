package main

import "fmt"

func emote(rating float64) string {
	if rating < 5 {
		return "bad"
	} else if 5 <= rating && rating < 7 {
		return "normal"
	} else if 7 <= rating && rating < 10 {
		return "good"
	} else {
		return "god"
	}
}

func main() {
	fmt.Println(emote(99))
}
