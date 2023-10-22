package snippets

import (
	"fmt"
)

func refFunc(i *int) {
	*i = 420
}

func refSlice(ii []int) {
	ii[0] = 1234
}

func refMap(iii map[int]int) {
	iii[1] = 1234
}

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and im walking")
}

func (d *dog) run() {
	d.first = "frajer"
	fmt.Println("My name is", d.first, "and im run")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func Pointers() {
	variable := 42
	address := &variable

	fmt.Println(variable)
	fmt.Println(address)
	fmt.Println(*address)
	fmt.Println(*&variable)

	*address = 43
	fmt.Println(variable)
	fmt.Println(address)
	fmt.Println(*address)

	fmt.Printf("%v %T\n", address, address)

	a := 1
	fmt.Println(a)
	refFunc(&a)
	fmt.Println(a)

	aa := []int{1, 2, 3, 4}
	fmt.Println(aa)
	refSlice(aa)
	fmt.Println(aa)

	aaa := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(aaa)
	refMap(aaa)
	fmt.Println(aaa)

	d1 := dog{"henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) -> ERROR run metoda ma pointer receiver

	d2 := &dog{"john"}
	d2.walk()
	d2.run()
	youngRun(d2)

	fmt.Println("pointers END------------")
}
