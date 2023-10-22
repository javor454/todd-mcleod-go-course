package snippets

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// go get golang.org/x/crypto/bcrypt

func Bcrypt() {
	pass := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(pass)
	fmt.Println(hash)

	//err = bcrypt.CompareHashAndPassword(hash, []byte(pass))
	err = bcrypt.CompareHashAndPassword(hash, []byte("wrong"))
	if err != nil {
		fmt.Println("spatny heslo buddy")
	} else {
		fmt.Println("gucci")
	}
	fmt.Println("bcrypt END------------")
}
