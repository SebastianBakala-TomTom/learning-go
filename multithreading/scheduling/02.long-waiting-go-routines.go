package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one phisical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg is used to manage concurency.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function.
	go func() {
		printPrime("A")
		wg.Done()
	}()

	// Create a goroutine from the uppercase function.
	go func() {
		printPrime("B")
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("\nWaiting To Finish")
	fmt.Println()
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime prints prime numbers for the first 5000 numbers
func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)

}
