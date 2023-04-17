package main

import "fmt"

// by value
func swap1(a int, b int) {
	var temp int = a
	a = b
	b = temp
}

// by address
func swap2(a *int, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

// go haven't pass by ref
// func swap3(&a int, &b int) {
// 	var temp int = a
// 	a = b
// 	b = temp
// }

// struct in go no need to use pointer
type course struct {
	name  string
	price float64
}

func discount(c *course) {
	c.price = c.price - 100
}

func discount2(c *course) {
	(*c).price = (*c).price - 100
}

type movie struct {
	name   string
	year   int
	rating float64
	genres []string
	voted  []float64
}

func (m *movie) addVote(r float64) {
	(*m).voted = append((*m).voted, r)

}

func main() {
	var i int = 64
	var p *int = &i
	fmt.Printf("i: %#v\n", i)
	fmt.Printf("&i: %#v\n", &i)
	fmt.Printf("p: %#v\n", p)
	fmt.Printf("&p: %#v\n", &p)
	fmt.Printf("*p: %#v\n", *p)
	a := 10
	b := 16
	swap1(a, b)
	fmt.Printf("a: %d b : %d\n", a, b)
	swap2(&a, &b)
	fmt.Printf("a: %d b : %d\n", a, b)

	// best practice
	// didnt good
	// c := course{"go",123}
	// d:=discount(&c)
	c := &course{"rust", 69}
	discount(c)
	fmt.Println("price ", c.price)

	// excercise
	m := &movie{
		name:   "god father",
		rating: 10,
		voted:  []float64{1, 2, 3, 4},
	}
	m.addVote(69)
	fmt.Printf("%#v", m.voted)
}
