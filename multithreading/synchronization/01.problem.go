// Data races problem
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines
var counter int

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

				// Capture the value of counter
				value := counter

				// Increment our local value of counter
				value++

				// This line, is a SysCall, which means that it causes the context switching.
				// Due to that, goroutines are swicthing it's context to the Waiting state before it updates the shared line (global variable).
				// This is the synchronization problem.
				// We can solve it by using atomic functions (good for small pieces of data 8bytes max) or mutex-es.
				fmt.Println(value)

				// Store the value back into counter.
				counter = value
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Printf("Final Counter: %d\n\n", counter)
}
