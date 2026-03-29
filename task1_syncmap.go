package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Store("key", val)
		}(i)
	}
	wg.Wait()

	value, _ := sm.Load("key")
	fmt.Printf("Value: %v\n", value)
}
