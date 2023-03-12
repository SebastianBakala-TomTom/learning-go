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
		lowercase()
		wg.Done()
	}()

	// Create a goroutine from the uppercase function.
	go func() {
		uppercase()
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("\nWaiting To Finish")
	fmt.Println()
	wg.Wait() // Wait() method is a DEMAND to wait for any remaining path of execution.

	// runtime.Gosched() // Never use Gosched() runtime method on production code.
	//It only a REQUESTS to execute any additional paths of execution. But there is no guarantee.

	fmt.Println("Terminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
	fmt.Println()
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
	fmt.Println()
}
