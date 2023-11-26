package ninja_excercises

import (
	"fmt"
)

type person struct {
}

func (p *person) speak() {
	fmt.Println("yoo")
}

type human interface {
	speak()
}

func speak(h human) {
	h.speak()
}

func Ex2() {
	p := person{}
	// speak(p) // NOK
	speak(&p) // OK
	p.speak() // OK
	fmt.Println("Ninja Exercise 2 END------------")
}
