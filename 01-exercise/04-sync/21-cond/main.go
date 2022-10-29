package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	c := sync.NewCond(&m)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()

		for len(sharedRsc) == 0 {
			time.Sleep(time.Second)
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"

	wg.Wait()
}
