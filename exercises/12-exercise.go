package exercises

import (
	"fmt"
	"time"
)

func TimedFunction(f func()) {
	start := time.Now()
	f()
	fmt.Printf("Function duration: %v\n", time.Since(start))
}

func Function() {
	time.Sleep(2 * time.Second)
	fmt.Println("Function done")
}

func Ex12() {
	TimedFunction(Function)
	fmt.Println("Exercise 12 END------------")
}
