package exercises

import (
	"fmt"
)

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}
func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

type person struct {
	first string
}

func changeV(p person, s string) person {
	p.first = s
	return p
}
func changeP(p *person, s string) {
	p.first = s
}

func Ex13() {
	aa := 1
	fmt.Println(&aa)
	var (
		a, b, c *string
		d       *int
	)

	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n

	fmt.Println(*a, *b, *c, *d)

	d1 := dog{"henry"}
	d1.walk()
	d1.run()
	//youngRun(d1)

	d2 := &dog{"padget"}
	d2.walk()
	d2.run()
	youngRun(d2)

	p1 := person{"john"}
	fmt.Println(p1)
	changeP(&p1, "1")
	fmt.Println(p1)
	p1 = changeV(p1, "2")
	fmt.Println(p1)

	//p2 := person{"frajer"}
	fmt.Println("Exercise 13 END------------")
}
