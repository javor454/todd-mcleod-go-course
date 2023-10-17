package exercises

import "fmt"

func Ex02() {
	var a string
	b := 11
	c, d := 11, 22
	var e float32 = 42.42
	g, h, _ := 11, 22, 33

	fmt.Printf("Zero value: %v %T\n", a, a)
	fmt.Printf("Shorthand declaration: %v %T\n", b, b)
	fmt.Printf("Multiple initializations: %v %T \t %v %T \n", c, c, d, d)
	fmt.Printf("Specificity: %v %T\n", e, e)
	fmt.Printf("Blank: %v %T \t %v %T\n", g, g, h, h)

	fmt.Println("exercise 02 END------------")
}
