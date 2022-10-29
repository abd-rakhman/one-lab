package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

var wg sync.WaitGroup

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	cores := runtime.GOMAXPROCS(runtime.NumCPU())
	var sum int64
	ch := make(chan int64)

	parts := (len(numbers) + cores - 1) / cores

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for val := range ch {
			sum += int64(val)
			i++
			if i == cores {
				return
			}
		}
	}()

	for i := 0; i < cores; i++ {
		go func(pos int) {
			ch <- Add(numbers[(pos * parts):((pos + 1) * parts)])
		}(i)
	}
	wg.Wait()

	return sum
}
