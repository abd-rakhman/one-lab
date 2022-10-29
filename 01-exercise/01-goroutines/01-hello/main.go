package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	fun("direct call")

	// defer wg.Done()

	go fun("first call")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("second call")
			time.Sleep(1 * time.Millisecond)
		}
	}()

	ff := fun

	go ff("third call")

	time.Sleep(5 * time.Second)

	fmt.Println("done..")
}
