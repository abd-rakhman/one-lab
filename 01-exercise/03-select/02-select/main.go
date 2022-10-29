package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
