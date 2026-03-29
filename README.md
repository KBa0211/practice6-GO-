## Task 2 Explanation
The final answer is not 1000 due to a race condition, as multiple goroutines concurrently read and overwrite the shared 'counter' variable without synchronization, leading to lost updates
