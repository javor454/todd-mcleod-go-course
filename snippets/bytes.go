package snippets

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Bytes() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	s := []byte("Hello gophers")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	b := bytes.NewBufferString("Cus kamo")
	fmt.Println(b.String())
	b.WriteString(" vitej ve svete go")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Po resetu")
	fmt.Println(b.String())
	b.Write([]byte(" zapsano jako pole bytu"))
	fmt.Println(b.String())

	fmt.Println("snippets bytes END------------")
}
