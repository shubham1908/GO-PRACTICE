package main

import (
	"fmt"
	"sync"
)

func threadUnsafeChannelSlice(n int) {
	var wg sync.WaitGroup
	ch := make(chan int)
	sliceA := []int{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(x int) {
			defer wg.Done()
			ch <- x
		}(i)
	}
	go func() {
		for v := range ch {
			// still unsafe because many goroutine may append to sliceA that can lead to race condition
			sliceA = append(sliceA, v) 
		}
	}()
	wg.Wait()
	close(ch)
	fmt.Printf("Added %d elements to slice, received %d\n", n, len(sliceA))
}
