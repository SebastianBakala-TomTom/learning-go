// Simple program to show how to access an exported identifier.

package main

import (
	"fmt"

	"github.com/SebastianBakala-TomTom/learning-go/decoupling/exporting/example_1/counters"
)

func main() {

	// Create a variable of the exported type and initialize the value of 10.
	counter := counters.AlterCounter(10)

	fmt.Printf("Counter: %s\n", counter)
}
