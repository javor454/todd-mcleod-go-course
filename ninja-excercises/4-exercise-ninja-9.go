package ninja_excercises

import (
	"fmt"
	"sync"
)

func Ex4() {
	gc := 200
	inc := 0
	var wc sync.WaitGroup
	wc.Add(gc)
	var mu sync.Mutex

	for i := 0; i < gc; i++ {
		go func() {
			mu.Lock()
			tmp := inc
			tmp++
			inc = tmp
			mu.Unlock()
			wc.Done()
		}()
	}
	wc.Wait()
	fmt.Println(inc)
	fmt.Println("Ninja Exercise 3 END------------")
}
