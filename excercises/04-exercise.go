package excercises

import "fmt"

type ByteSize int

const (
	_           = iota // ignore first value
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func Ex04() {
	fmt.Printf("Decimal: %d \t Binary: %b\n", KB, KB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", MB, MB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", GB, GB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", TB, TB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", PB, PB)
	fmt.Printf("Decimal: %d \t Binary: %b\n", EB, EB)
	fmt.Println("exercise 4 END------------")
}
