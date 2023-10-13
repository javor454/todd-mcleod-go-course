package excercises

import (
	"fmt"
	"math"
)

func Ex06() {
	a, b, c := "yo", 1, 232.32132
	fmt.Printf("%s %T\n%d %T\n%f %T\n", a, a, b, b, c, c)

	var d int8 = math.MaxInt8
	var e uint8 = math.MaxUint8
	fmt.Println(d, e)
	fmt.Println("exercise 6 END------------")
}
