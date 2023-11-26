package ninja_excercises

import (
	"fmt"
	"runtime"
	"sync"
)

func Ex3() {
	gc := 200
	inc := 0
	var wc sync.WaitGroup
	wc.Add(gc)

	for i := 0; i < gc; i++ {
		go func() {
			tmp := inc
			runtime.Gosched()
			tmp++ // creates race condition as multiple goruts are changing state
			inc = tmp
			wc.Done()
		}()
	}
	wc.Wait()
	fmt.Println("Ninja Exercise 3 END------------")
}
