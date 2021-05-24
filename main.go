package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	gs := 1000
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			locCount := count
			runtime.Gosched()
			locCount++
			count = locCount
			fmt.Println("current count", count)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}
