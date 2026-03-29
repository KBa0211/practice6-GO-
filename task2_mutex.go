package main

import (
	"fmt"
	"sync"
)

// Explanation for teamlead: The final answer is not 1000 due to a race condition, as multiple goroutines concurrently read and overwrite the shared 'counter' variable without synchronization, leading to lost updates.
func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ { // [cite: 126]
		wg.Add(1)   // [cite: 127]
		go func() { // [cite: 128]
			defer wg.Done() // [cite: 129]
			mu.Lock()
			counter++ // [cite: 130]
			mu.Unlock()
		}() // [cite: 131]
	} // [cite: 132]
	wg.Wait()            // [cite: 133]
	fmt.Println(counter) // [cite: 134]
}
