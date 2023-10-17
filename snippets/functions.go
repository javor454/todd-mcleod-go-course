package snippets

import (
	"fmt"
	"log"
	"strconv"
)

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

func foo() {
	fmt.Println("cus")
}
func bar(s string) {
	fmt.Println("cus", s)
}
func aloha(s string) string {
	return fmt.Sprint("cus ", s)
}
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " ma tolik let:"), age
}
func variadic(name string, i ...int) {
	fmt.Println(i)
	fmt.Println(name)
	fmt.Printf("typ variadic %T\n", i)
}
func frajer() {
	fmt.Println("defer frajer funkce")
}
func buddy() {
	fmt.Println("buddy")
}

func Functions() {
	foo()
	bar("frajirek")
	fmt.Println(aloha("kamos"))
	fmt.Println(dogYears("alik", 3))

	variadic("kkkk", 1, 2, 3, 4, 5)
	variadic("empty")
	variadic("unfurling slice", []int{1, 2, 3}...)

	defer frajer() // odlož provolání funkce na dokončení upper funkce
	buddy()

	metody()
	fmt.Println("Functions end ----------")
	anonFuncs()
	funcExpression()
	returnFunc()
	callbackFunc()
	closureFunc()
	fmt.Println(factorial(4))
	fmt.Println(factorialLoop(4))
}

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	licenceToKill bool
}
type count int

func (p person) predstavSe() {
	fmt.Printf("jsem %v %v a je mi %v let\n", p.first, p.last, p.age)
}
func (sa secretAgent) predstavSe() {
	fmt.Printf("jsem tajny agent %v %v a je mi %v let, licence na zabijeni %v\n", sa.first, sa.last, sa.age, sa.licenceToKill)
}
func (c count) String() string {
	return fmt.Sprintf("jsem pocet %v", strconv.Itoa(int(c)))
}

func (p person) String() string {
	return fmt.Sprintf("jsem %v %v a je mi %v let\n", p.first, p.last, p.age)
}
func (sa secretAgent) String() string {
	return fmt.Sprintf("jsem tajny agent %v %v a je mi %v let, licence na zabijeni %v\n", sa.first, sa.last, sa.age, sa.licenceToKill)
}

func wrapperLogInfo(s fmt.Stringer) {
	log.Println("LOG FUNKCI:", s.String())
}

func rekniNeco(h human) {
	h.predstavSe()
}

type human interface {
	predstavSe()
}

func metody() {
	p1 := person{
		"karel",
		"novak",
		23,
	}
	p1.predstavSe()
	sa1 := secretAgent{
		p1,
		true,
	}
	sa1.predstavSe()
	c1 := count(13)
	fmt.Println(c1)

	rekniNeco(p1)
	rekniNeco(sa1)

	fmt.Println(p1)
	fmt.Println(sa1)

	log.Println(p1)
	log.Println(sa1)
	wrapperLogInfo(p1)
	wrapperLogInfo(sa1)

	s := "kamos"
	bs := []byte(s)
	sbs := string(bs)
	fmt.Printf("%v %T %v %T %v %T\n", s, s, bs, bs, sbs, sbs)
}

// signature anon funkce func (p parameter(s)) (r return(s)) { code }
func anonFuncs() {
	func() {
		fmt.Println("anon func")
	}()
	func(s string) {
		fmt.Println("hello", s)
	}("franta")
}

func funcExpression() {
	x := func(s string) string {
		return s
	}

	x("cus")

	y := func(s string) string {
		return s
	}("cus")
	fmt.Println(y)

	z := func(s string) {
		fmt.Println("cus", s)
	}
	z("kamo")
}

func funcReturnFunc() func() int {
	return func() int {
		return 42
	}
}

func funcA() int {
	return 43
}

func returnFunc() {
	b := funcA()
	fmt.Printf("%v %T\n", b, b)

	a := funcReturnFunc()
	fmt.Printf("%v %T\n", a(), a)
}

func add(a, b int) int {
	return a + b
}
func substract(a, b int) int {
	return a - b
}

func doOperation(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func callbackFunc() {
	fmt.Println(doOperation(25, 25, add))
	fmt.Println(doOperation(23, 11, substract))
}

func closureFunc() {
	f := incrementor()
	fmt.Println(f(), f(), f())

	g := incrementor() // jina hodnota
	fmt.Println(g(), g(), g())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func factorial(i int) int {
	fmt.Println("i=", i)
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func factorialLoop(i int) int {
	total := 1
	for i > 0 {
		fmt.Println("i=", i)
		total = total * i
		i--
	}
	return total
}
