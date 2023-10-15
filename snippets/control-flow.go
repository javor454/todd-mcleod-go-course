package snippets

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("init controlflow runs before anything in main ")
}
func ControlFlow() {
	fmt.Println("start controlflow func")

	// podminky
	x := 1
	if x > 0 {
		fmt.Println("vetsi")
	} else if x < 0 {
		fmt.Println("mensi")
	} else {
		fmt.Println("rovny")
	}

	y := x + 1
	if y < 0 && x < 0 {
		fmt.Println("oba mensi")
	}

	if y > 0 || x > 0 {
		fmt.Println("alespon jeden vetsi")
	}

	if y != 0 {
		fmt.Println("neni nula")
	}

	// statement idiom omezi scope promenne z pouze na podminku => cil je omezit scope co nejvice
	if z := rand.Intn(2) * x; z > 1 {
		fmt.Printf("statement idiom promenna vetsi nule z=%v\n", z)
	} else {
		fmt.Printf("statement idion promenna mensi nebo rovna nule z=%v\n", z)
	}

	// switch
	a := rand.Intn(3)
	switch {
	case a == 0:
		fmt.Println("nula")
	default:
		fmt.Println("nenula")
	}

	switch a {
	case 0:
		fmt.Println("nula")
	default:
		fmt.Println("nenula")
	}

	selectStatement()
	forCykly()

	fmt.Println(83 % 40) // zbytek
}

func selectStatement() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Printf("po kanalu 1 prisla zprava (%v)\n", v1)
	case v2 := <-ch2:
		fmt.Printf("po kanalu 2 prisla zprava (%v)\n", v2)
	}
}

func forCykly() {
	for a := 1; a < 5; a++ {
		fmt.Println(a)
	}

	b := 1
	for b < 5 {
		fmt.Println(b)
		b++
	}

	c := 1
	for {
		if c < 5 {
			fmt.Println(c)
			c++
			continue
		}
		break
	}

	for i, v := range []int{1, 2, 3} {
		fmt.Printf("index = %v value = %v\n", i, v)
	}

	for k, v := range map[string]int{
		"Ahoj": 1,
		"Cus":  2,
	} {
		fmt.Printf("klic = %v hodnota = %v\n", k, v)
	}

	for k := range []int{1, 2, 3} {
		fmt.Printf("klic = %v\n", k)
	}

	for _, v := range []int{1, 2, 3} {
		fmt.Printf("klic se nepouziva, hodnota = %v\n", v)
	}

	for pos, char := range "ahoj" {
		fmt.Printf("pos = %d char = %#U\n", pos, char)
	}

	mapik := map[string]int{
		"jirka": 28,
	}
	if hodnota, ok := mapik["jirka"]; ok {
		fmt.Printf("hodnota = %v\n", hodnota)
	} else {
		fmt.Printf("oops, vracim nulovou hodnotu intu = %v\n", hodnota)
	}

	if _, ok := mapik["jirka"]; ok {
		fmt.Println("mapa obsahuje klic")
	}
}
