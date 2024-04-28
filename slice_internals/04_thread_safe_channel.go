package main

import (
	"fmt"
	"sync"
)

func threadSafeChannelSlice(n int) {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int)
	sliceA := []int{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(x int) {
			defer wg.Done()
			ch <- x
		}(i)
	}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			// still unsafe because many goroutine may append to sliceA 
			// that can lead to race condition
			mu.Lock()
			sliceA = append(sliceA, v)
			mu.Unlock()
		}
	}()
	wg.Wait()
	close(ch)
	wg2.Wait()
	fmt.Printf("Added %d elements to slice, received %d\n", n, len(sliceA))
	// sort.Ints(sliceA)
	fmt.Println(len(sliceA))

	// wg2 is important because there might be a chance that 
	// go routine which is appending element to the slice is taking time 
	// (maybe because of locks)
	// due to which although eventually the slice will contain correct elements 
	// but when compiler is prinitng the size of slice, the element maynot have 
	// added to the slice

}
