package excercises

import (
	"fmt"
	"math/rand"
)

func Ex05() {
	x := rand.Intn(250)
	fmt.Printf("x = %v\n", x)
	switch {
	case x < 100:
		fmt.Println("mezi 0 a 100")
	case x >= 101 && x <= 200:
		fmt.Println("mezi 101 a 200")
	case x >= 201 && x <= 250:
		fmt.Println("mezi 201 a 250")
	default:
		fmt.Println("vetsi nez 250")
	}
	if x <= 100 {
		fmt.Println("mezi 0 a 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("mezi 101 a 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("mezi 201 a 250")
	} else {
		fmt.Println("vetsi nez 250")
	}

	b := 0
	for b < 101 {
		fmt.Printf("iterace %v\n", b)
		for a := 0; a < 100; a++ {
			fmt.Printf("%v ", a)
		}
		fmt.Printf("\n")
		b++
	}

	for a := 1; a < 43; a++ {
		fmt.Printf("iterace %v\n", a)
		x := rand.Intn(5)
		switch {
		case x < 3:
			fmt.Println("mensi nez 3")
		case x > 3:
			fmt.Println("vetsi nez 3")
		default:
			fmt.Println("tri")
		}
	}

	for i, v := range []int{1, 2, 3, 4} {
		fmt.Printf("index %v value %v\n", i, v)
	}
	for k, v := range map[string]int{"a": 1, "b": 2} {
		fmt.Printf("key %v value %v\n", k, v)
	}
	xi := []int{1, 2}
	fmt.Printf("slice hodnota %v\n", xi[0])
	mapik := map[string]int{"a": 123}
	fmt.Printf("map hodnota %v\n", mapik["a"])

	if x := rand.Intn(5); x == 3 {
		fmt.Println("x je 3")
	}

	fmt.Println(true && true, true && false, true || true, true || false, !true)
	fmt.Println("exercise 05 END------------")
}
