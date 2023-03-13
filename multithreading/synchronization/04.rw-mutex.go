// Data races problem
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// data is a slice that will be shared
var data []string

// rwMutex is used to define a critical section of code.
var rwMutex sync.RWMutex

// Number of reads occuring at any given time.
var readCount int64

// init is called prior to main
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Number of goroutines to use
	const grs = 1

	// wg is used to manage concurency
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create a writer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(rand.Int())
		}
		wg.Done()
	}()

	// Create eight reader goroutines
	for i := 0; i < 800; i++ {
		go func(id int) {
			for {
				reader(id)
			}
		}(i)
	}
	// Wait for the goroutines to finish
	wg.Wait()

}

func writer(id int) {

	// Only allow one goroutine to read/write to the slice at a time.
	rwMutex.Lock()
	{
		// Capture the current read count.
		// Keep this safe though due we can without this call
		rc := atomic.LoadInt64(&readCount)

		// Perform some work since we have a full lock.
		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", id))

	}
	rwMutex.Unlock()
	// Release the lock

}

func reader(id int) {

	// Any goroutine can read when no write operation is taking place.
	rwMutex.RLock()
	{
		// Increment the read count value by 1.
		rc := atomic.AddInt64(&readCount, 1)

		// Perform some read work and display values.
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)

		// Decrement the readCount value by 1.
		atomic.AddInt64(&readCount, -1)

	}
	rwMutex.RUnlock()
	// Release the read-lock.
}
