package exercises

import (
	"fmt"
	"log"
	"os"
)

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file '%v', error: %v\n", fileName, err)
	}

	return xb, nil
}

func Ex08() {
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatalf("error in main: %v\n", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

	fmt.Println("Exercise 08 END------------")
}
