package snippets

import (
	"fmt"
	"runtime"
	"sync"
)

func Mutex() {
	fmt.Println("CPU:", runtime.NumCPU())
	count := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GS:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println("Exercise 16 END------------")
}
