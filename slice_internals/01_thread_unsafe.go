package main

import (
	"fmt"
	"sync"
)

func threadUnsafeSlice(n int) {
	var wg sync.WaitGroup
	sliceA := []int{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(x int) {
			defer wg.Done()
			sliceA = append(sliceA, x)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Added %d elements to slice, received %d\n", n, len(sliceA))
}
