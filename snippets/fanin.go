package snippets

import "fmt"

func send(c <-chan int) {
	fmt.Println(<-c)
}

func receive(c chan<- int) {
	c <- 1
}

func Channels() {
	c := make(chan int)
	cr := make(chan<- int) // receive
	cs := make(<-chan int) // send
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	ch := make(chan int)

	go send(ch) // make bidirectional channel more specific -> send channel
	receive(ch) // make bidirectional channel more specific -> receive channel
	fmt.Println("Channel end ----------")
}
