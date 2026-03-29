package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func main() {
	safeMap := SafeMap{m: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			safeMap.mu.Lock()
			safeMap.m["key"] = val
			safeMap.mu.Unlock()
		}(i)
	}
	wg.Wait()

	safeMap.mu.RLock()
	value := safeMap.m["key"]
	safeMap.mu.RUnlock()
	fmt.Printf("Value: %d\n", value)
}
