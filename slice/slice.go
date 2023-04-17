package main

import "fmt"

func main() {
	// fmt.Println("Hello")
	// slice like vector in c++
	// 1) dynamic
	// 2) if null cannot assign  , array must append first
	// 3) it copy box then append after that point it
	// 4) like math slice is [)
	// 			[2,6] ->   2 - (6-2) -> 2, 3, 4, 5
	// 5) make -> intial value follow type of slice

	var slice = []string{"rust", "go", "huskell"}
	fmt.Printf("%#v\n ", slice)
	fmt.Printf("%#v\n ", slice[0])

	slice2 := append(slice, "typescript")
	fmt.Printf("%#v\n ", slice2)
	slice = append(slice, "Elixr")
	fmt.Printf("%#v\n ", slice)
	slice3 := append(slice, "Java", "C++")
	fmt.Printf("%#v\n ", slice3)

	// slicing
	begin := slice[0:]
	fmt.Printf("%#v\n ", begin)
	all := slice[:]
	fmt.Printf("%#v\n ", all)

	// make
	make := make([]int, 3)
	fmt.Printf("%#v\n", make)

	// excercise
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}
	// vote := xs
	// for i =: 0; i< len(ys) ; i++ {
	// 	vote = append(ys, ys[i])
	// }
	vote := append(xs, ys...)
	fmt.Printf("%#v\n", vote)
	vote = append(vote, 1)
	fmt.Printf("%#v\n", vote)

}
