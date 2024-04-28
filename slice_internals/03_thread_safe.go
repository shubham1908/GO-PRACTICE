package main

import (
	"fmt"
	"sync"
)

func threadSafeSlice(n int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sliceA := []int{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(x int) {
			defer wg.Done()
			mu.Lock()
			sliceA = append(sliceA, x)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Added %d elements to slice, received %d\n", n, len(sliceA))
}
