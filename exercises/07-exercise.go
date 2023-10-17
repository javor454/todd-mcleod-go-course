package exercises

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person1 struct {
	first string
}

func (p person1) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func Ex07() {
	p := person1{
		"Karel",
	}

	f, err := os.Create("person.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	var b bytes.Buffer

	_ = p.writeOut(f)
	_ = p.writeOut(&b)

	fmt.Println(b.String())

	fmt.Println("Exercise 07 END------------")
}
