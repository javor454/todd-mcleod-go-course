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

	fmt.Println("pointers END------------")
}
