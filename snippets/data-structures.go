package snippets

import "fmt"

func DataStructures() {
	array()
	slice()
}

func array() {
	// array literal
	ai := [3]int{1, 2, 3} // nazev promenne ai = array of ints
	fmt.Println(ai)

	// infer velikosti pole
	bs := [...]string{"cus", "brasko"}
	fmt.Println(bs)

	var ci [2]int
	ci[0] = 1
	//ci[1] = 2 // doplni nulovou hodnotu
	fmt.Println(ci)
	fmt.Printf("pole c ma hodnoty %v a velikost %v a typ %T\n", ci, len(ci), ci) // typ = [2]int - velikost pole je soucasti definice
}

func slice() {
	xs := []string{"cus", "buddy"}
	fmt.Println(xs, len(xs))
	for k, v := range xs {
		fmt.Println(k, v)
	}
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println(xs[0])
	//fmt.Println(xs[10]) // error index out of range

	// append to slice
	xs = append(xs, "yoo", "cus cus") // buildin function, variadic parametr = prijima dynamicky pocet hodnot stejneho typu
	fmt.Println(xs)

	// slice a slice
	fmt.Println(xs[1:2])
	fmt.Println(xs[2:])
	fmt.Println(xs[:2])
	fmt.Println(xs[:])

	// delete from slice
	fmt.Println(append(xs[:1], xs[2:]...))

	// vytvoreni pomoci make
	xs2 := make([]int, 0, 10) // delka 0 kapacita 10
	xs2 = append(xs2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("hodnoty %v typ %T delka %v kapacita %v\n", xs2, xs2, len(xs2), cap(xs2)) // delka 10 kapacita 10
	xs2 = append(xs2, 10, 11, 12)
	fmt.Printf("hodnoty %v typ %T delka %v kapacita %v\n", xs2, xs2, len(xs2), cap(xs2)) // delka 13 kapacita 20

	// TODO: 13 - Grouping data values - arrays & slice: 012 Slice - multidimensional slice
}
