package snippets

import (
	"fmt"
	"slices"
)

func DataStructures() {
	array()
	slice()
	mapy()
	structs()
	fmt.Println("DataStructures end ----------") // zbytek
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

	// multidimenzionalni slice
	xa := []string{"kuba", "pivko"}
	xb := []string{"jirka", "cigo"}
	xc := [][]string{xa, xb}
	fmt.Println(xc)

	// oba slicy maji pointer na stejne interni pole
	xd := []int{1, 2, 3}
	xe := xd
	fmt.Println(xd, xe)
	xd[0] = 123
	fmt.Println(xd, xe)

	// kazdy slice ma vlastni interni pole
	xf := []int{1, 2, 3}
	xg := make([]int, 3)
	fmt.Println(xf, xg)
	copy(xg, xf)
	fmt.Println(xf, xg)
	xf[0] = 123
	fmt.Println(xf, xg)

	//
	xh := []int{2, 3, 1}
	fmt.Println(xh)
	sortModify(xh)
	fmt.Println(xh)

	xch := []int{2, 3, 1}
	fmt.Println(xch)
	sortNoModify(xch)
	fmt.Println(xch)

}

func mapy() {
	am := map[string]int{
		"Jirka": 11,
		"Tomas": 22,
		"Karel": 33,
	}
	fmt.Println(am["Jirka"])
	fmt.Println(am)

	an := make(map[string]int)
	an["Lukas"] = 11
	an["Frajer"] = 34

	fmt.Println(len(an))

	for k, v := range am {
		fmt.Println(k, v)
	}
	for k := range am {
		fmt.Println(k)
	}
	for _, v := range am {
		fmt.Println(v)
	}

	fmt.Println(an)
	delete(an, "Frajer")
	fmt.Println(an)
	delete(an, "xxx")      // no panic
	fmt.Println(an["xxx"]) // zero value
	if v, ok := an["Lukas"]; ok {
		fmt.Printf("ok value = %v\n", v)
	} else {
		fmt.Println("missing")
	}

	m := map[string]int{
		"a": 0,
	}
	fmt.Println(m)
	m["a"]++
	fmt.Println(m)
	m["a"]++
	fmt.Println(m)
}

func structs() {
	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{
		first: "james",
		last:  "bond",
		age:   42,
	}

	fmt.Printf("%T %v\n", p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)

	type secretAgent struct {
		person
		licenceToKill bool
		first         string
		// last - inner type promotion -> secretAgent.last
		// age - inner type promotion -> secretAgent.age
	}

	sa1 := secretAgent{p1, true, "frajer"}
	fmt.Printf("%T %#v\n", sa1, sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.person.first, sa1.licenceToKill)

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "kamos",
		last:  "novak",
		age:   11,
	}
	fmt.Printf("%T %#v\n", p2, p2)
	fmt.Println(p2.first, p2.last, p2.age)

	// do pojmenovaneho typu muzu priradit kompatibilni anonymni typ
	a1 := struct {
		a string
	}{
		a: "Karel",
	}
	type example struct {
		a string
	}
	a2 := example{
		a: "Jarda",
	}
	fmt.Println(a1, a2)
	a2 = a1
}

func sortModify(a []int) {
	slices.Sort(a)
}

func sortNoModify(a []int) {
	b := make([]int, len(a))
	copy(b, a)
	slices.Sort(b)
}
