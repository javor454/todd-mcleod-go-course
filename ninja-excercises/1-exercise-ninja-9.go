package ninja_excercises

import (
	"fmt"
	"sync"
)

func Ex1() {
	var wg sync.WaitGroup
	gs := 2
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func(index int) {
			fmt.Printf("g:%v\n", index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Ninja Exercise 1 END------------")
}
