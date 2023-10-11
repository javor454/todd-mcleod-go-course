// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import "fmt"

func main() {
	fmt.Println("hello 👈👈👈")

	const koko = "ahoj"
	fmt.Printf("%v buddy pls", koko)
	fmt.Println(`
		ahoj
		buddy
	`)

	// variable can vary
	var h int = 44

	// infer int
	g := 45

	// best practice - use shorthand
	j, k, l := 1, 2, 3

	// unused vars throw error
	// m := "asd"

	// zero value pro typ - int = 0, bool = false
	var ab int
	var ac bool

	number := 42
	fmt.Printf("original: %v, binary: %b, hexa: %X", number, number, number)

	// v go neni casteni ale konverze
	fmt.Println(h, g, j, k, l, ab, ac)

	yy := 42
	zz := 42.0
	fmt.Printf("%v je typu %T\n", yy, yy)
	fmt.Printf("%v je typu %T\n", zz, zz)

	var yyy float32 = 12.234
	// yyy = zz // error: nekompatibilni assignment


	zz = float64(yyy)
	fmt.Println(zz, yyy)

	// TODO lekce 04 - 004 Values, types, conversion, scope, & housekeeping
}
