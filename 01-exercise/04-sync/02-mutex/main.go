package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	deposit := func(amount int) {
		mu.RLock()
		balance += amount
		mu.RUnlock()
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
