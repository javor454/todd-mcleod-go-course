package snippets

import "fmt"

// package block scope -> dostupne v packagich main
// x := 32 // variable shadowing

func Scope() {
	// block scope

	x := 33
	fmt.Println(x)
}
