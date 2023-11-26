package snippets

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func fce1() {
	for i := 0; i < 5; i++ {
		fmt.Println("fce 1:", i)
	}
}

func fce2() {
	for i := 0; i < 5; i++ {
		fmt.Println("fce 2:", i)
	}
}

func Concurency() {
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.NumGoroutine())
	unbufferedChannel()
	bufferedChannel()
	waitGroupa()
	// fce main je prvni gorutina
	// nyni nastartuji novou gorutinu, nez se nastartuje tak vsak main dokonci sekvenci ukolu a exitne -> vsechny gorutiny skonci
	go fce1()
	fce2()

	// TODO: 26 Concurency: 004 Documentation
	fmt.Println("concurency END------------")
}

func unbufferedChannel() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second) // simulace prace
		ch <- 5                 // poslu kanalu hodnotu v gorutine abych neblokoval hlavni gorutinu
	}()
	x := <-ch // hlavni gorutina zde blokuje a ceka na hodnotu z kanalu
	fmt.Println(x)
}

func bufferedChannel() {
	ch := make(chan int, 2) // kapacita 2
	ch <- 1                 // nemusi byt v nove gorutine kvuli bufferu 2
	ch <- 2

	fmt.Println(<-ch) // 1 - dostavam hodnoty v poradi ve kterem jsem je poslal
	fmt.Println(<-ch) // 2
}

func waitGroupa() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // increment counter of awaited goroutines
		go func() {
			defer wg.Done()                                         // when done decrement counter
			time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second) // simulate work
			fmt.Println("gorutina done")
		}()
	}
	wg.Wait()
	fmt.Println("jobs done")
}
