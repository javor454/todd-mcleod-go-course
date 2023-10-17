package snippets

import "fmt"

// package pouzivat poslednim slovem importu

// func add(x, y int) int { // oba argumenty maji typ int
//
//		return x + y
//	}
func add2(x bool, y int) {
	fmt.Printf("hodnota %v ma typ %T, hodnota %v ma typ %T\n", x, x, y, y)
}
func swap(first, second string) (string, string) {
	return second, first
}
func swap2(first, second string) (x, y string) {
	x = second
	y = first
	return
}

const Borec = "Frajer"

func Starter() {
	fmt.Println("hello 👈👈👈")

	const koko = "ahoj"
	fmt.Printf("%v buddy pls", koko)
	fmt.Println(`
		ahoj
		buddy
	`)

	// variable can vary
	var h int = 44

	// infer int, shorthand declarace je mozne pouzit pouze ve funkci
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

	// primitive types = poskytuje lagys jako zakladni block
	// aggregate types (kombinuje nekolik elementu)
	//   array, slice, struct = drzi hodnoty stejnych typu
	// composite type
	//   struct = drzi hodnoty ruznych typu

	// exported (dostupny) x unexported (nedostupny) z package

	fmt.Println(add(1, 2))
	add2(true, 2)
	fmt.Println(swap("prvni", "druhy"))
	fmt.Println(swap2("jedna", "dva"))

	// 1, 2  = initializer
	var aaa, bbb int = 1, 2
	var aaaa, bbbb = 1, 2
	aaaaa, bbbbb := 1, 2
	fmt.Println(aaa, bbb, aaaa, bbbb, aaaaa, bbbbb)

	// BEST PRACTICE: pokud potrebuju integer pouziju INT, neco jineho jen v pripade ze chci mit specifickou velikost nebo aby byl unsigned
	iii := 42
	fff := float64(iii)
	uuu := uint(fff)
	fmt.Println(uuu)

	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day", Borec)

	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1) // posouvam bit o jeden doleva

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)
}
