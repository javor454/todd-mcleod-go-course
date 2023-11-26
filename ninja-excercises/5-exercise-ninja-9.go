package ninja_excercises

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Ex5() {
	gc := 200
	var inc int64
	var wc sync.WaitGroup
	wc.Add(gc)

	for i := 0; i < gc; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			wc.Done()
		}()
	}
	wc.Wait()
	fmt.Println(inc)
	fmt.Println("Ninja Exercise 3 END------------")
}
