// Simple program to show how to access an exported identifier.

package main

import (
	"fmt"

	"github.com/SebastianBakala-TomTom/learning-go/decoupling/exporting/example_3/counters"
)

func main() {

	// Create a variable of the exported type and initialize the value of 10.
	// Do not use this kind of syntax. It's annoying to use
	counter := counters.New(10)

	fmt.Printf("Counter: %s\n", counter)
}
