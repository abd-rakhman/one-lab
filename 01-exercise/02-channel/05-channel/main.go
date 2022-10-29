package main

import "fmt"

func main() {
	// TODO: create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	owner := func() chan int {
		ch := make(chan int)
		// defer close(ch)

		// for i := 0; i < 5; i++ {
		// 	ch <- i
		// }
		ch <- 1
		return ch
	}

	// consumer := func(ch <-chan int) {
	// 	// read values from channel
	// 	for v := range ch {
	// 		fmt.Printf("Received: %d\n", v)
	// 	}
	// 	fmt.Println("Done receiving!")
	// }

	ch := owner()
	fmt.Println(<-ch)
	// close(ch)
	// consumer(ch)
}