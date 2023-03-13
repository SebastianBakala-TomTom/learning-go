// Data races problem
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines
var counter int

// mutex is used to define a critical section of code.
// Here defined as a global variable. We don't do that. It's only for educational purposes.
// Most of time it should be a field in a struct.
// When mutex is a field in a struct, this struct can no longer be copied. Copy of mutex creates a new value of mutex.
var mutex sync.Mutex

func main() {

	// Number of goroutines to use
	const grs = 2

	// wg is used to manage concurency
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {

				// Only allow one goroutine through this critical lines of code at a time.
				mutex.Lock()
				{
					// Capture the value of counter
					value := counter

					// Increment our local value of counter
					value++

					// Using printing functions inside the mutex is a waste of the latency and backpressure.
					// fmt.Println(value)
					// That's why you should do only and only what you really need inside the mutex.

					// Store the value back into counter.
					counter = value
				}
				mutex.Unlock()
				// Release the lock and allow any waiting goroutine through.
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Printf("Final Counter: %d\n\n", counter)
}
