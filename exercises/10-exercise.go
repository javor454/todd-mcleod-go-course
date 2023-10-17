package exercises

import (
	"fmt"
	"math"
)

func foo() int {
	return 1
}
func bar() (int, string) {
	return 1, "cus"
}
func foo2(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
func bar2(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func (p person2) speak() {
	fmt.Printf("moje jmeno je %v a vek je %v\n", p.first, p.age)
}

type person2 struct {
	first string
	age   int
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.width * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(math.Pi*c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func Ex10() {
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(foo2([]int{1, 2, 3}...))
	fmt.Println(bar2([]int{1, 2, 3}))

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	persona := person2{
		first: "karel",
		age:   11,
	}
	persona.speak()

	info(circle{radius: 1.23})
	info(square{
		length: 1,
		width:  2,
	})

	fmt.Println("Exercise 10 END------------")
}
