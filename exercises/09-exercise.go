package exercises

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type animal struct {
	name string
	age  int
}

func (a animal) makeNoise() {
	fmt.Println("rawr")
}
func (a animal) String() string {
	return fmt.Sprintf("Zvire %v je stare %v let.\n", a.name, a.age)
}
func prettyLog(s string) {
	log.Println("Sexy log:", s)
}
func funFunc() func() int {
	return func() int {
		return 11
	}
}

func funFunc2(f func(i int) int) int {
	return f(12)
}

func closure() int {
	x := 0
	return func(i int) int {
		for x < 5 {
			i += x
			x++
		}
		return i
	}(5)
}

func factorial(i int) int {
	x := 1
	fmt.Println("i =", i)
	if i > x {
		return i * factorial(i-1)
	}

	return i
}

func Ex09() {
	variadic([]int{1, 2, 3}...)
	defer fmt.Println("odlozeno na konec")

	a := animal{
		"micka",
		12,
	}
	a.makeNoise()
	fmt.Println(a.String())

	prettyLog("loguji progress")

	filename := "test.txt"
	writer(filename, "text na zapsani do souboru")
	fmt.Println(reader(filename))

	buffer := bytes.NewBufferString("cuus")
	buffer.WriteString(" brasko")
	buffer.Write([]byte(" jak to jde?"))
	fmt.Println(buffer.String())

	func(s string) {
		fmt.Println(s)
	}("anonymni funkce")

	ab := func(i int) int {
		return i + 123
	}
	fmt.Println(ab(132))

	fmt.Println(funFunc()())

	fmt.Println(funFunc2(func(i int) int {
		return i / 2
	}))

	fmt.Println(closure())

	fmt.Println(factorial(13))
	fmt.Println("Exercise 09 END------------")
}

func reader(filename string) string {
	xb, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("problem s ctenim souboru %v\n", err)
	}

	return string(xb)
}

func writer(filename, content string) {
	f, err := os.Create(filename)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		log.Fatalf("problem s vytvorenim souboru %v\n", err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalf("problem se zapisem %v\n", err)
	}
}

func variadic(i ...int) {
	fmt.Println(i)
}
