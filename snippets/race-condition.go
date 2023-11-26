package snippets

import (
	"fmt"
	"runtime"
	"sync"
)

func RaceCondition() {
	fmt.Println("CPU:", runtime.NumCPU())
	count := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// scope v je omezeny - kazda gorutina ma svuj
			// vsechny gorutiny vsak pristupuji ke stejnemu countu
			v := count
			// time.Sleep(time.Second * 2)
			runtime.Gosched() // go run something else
			v++
			count = v // dojde k race condition
			wg.Done()
		}()
		fmt.Println("GS:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println("Exercise 15 END------------")
}
