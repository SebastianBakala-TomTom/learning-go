// Simple program to show how to access an exported identifier.

package main

import (
	"fmt"
)

func main() {

	// Create a variable of the exported type and initialize the value of 10.
	counter := counters.alterCounter(10)

	// ./example.go:14: cannot refer to unexported name counters.alterCounter
	// ./example.go:14: undefined: counters.alterCounter

	fmt.Printf("Counter: %s\n", counter)
}
