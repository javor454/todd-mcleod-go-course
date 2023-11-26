package snippets

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Atomic() {
	fmt.Println("CPU:", runtime.NumCPU())
	var count int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println("counter:", atomic.LoadInt64(&count))
			wg.Done()
		}()
		fmt.Println("GS:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println("Exercise 17 END------------")
}
