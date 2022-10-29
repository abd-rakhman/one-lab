package main

import (
	"fmt"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- int) {
	ch1 <- 1
}

func relayMsg(ch1 <-chan int, ch2 chan<- int) {
	// recv message on ch1
	msg := <-ch1
	// send it on ch2
	ch2 <- msg
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan int)
	ch2 := make(chan int)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// recv message on ch2
	fmt.Printf("message is: %d\n", <-ch2)
}
