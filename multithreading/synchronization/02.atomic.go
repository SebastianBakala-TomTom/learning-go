// Data races problem
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all goroutines
// Because we are using the 'atomic' package, we need to use precision based intigers.
var counter int64

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
				// This is the fastest way to ensure synchonization between multiple threads.
				atomic.AddInt64(&counter, 1)
				// But it cuts us off from possibility to implement a custom logic like in 01.problem.go file.
				// Here comes the mutexes with their help.
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Printf("Final Counter: %d\n\n", counter)
}
