package snippets

import "fmt"

func sendVals(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // if we dont close, range will block and wait for more values
}

func ChannelsRange() {
	ch := make(chan int)
	go sendVals(ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("ChannelRange end ----------")
}
