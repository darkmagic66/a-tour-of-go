package main

import "fmt"

type course struct {
	name  string
	price float64
}

type movie struct {
	name   string
	year   int
	rating float64
	genres []string
}

func (c course) discount(discount float64) float64 {
	p := c.price - discount
	return p

}

func main() {
	fmt.Println("go")
	c := course{
		price: 69,
	}
	fmt.Printf("%#v\n", c)
	// excercise
	gf := movie{
		name: "god father",
		year: 1972,
	}
	fmt.Printf("%#v", gf)
}
