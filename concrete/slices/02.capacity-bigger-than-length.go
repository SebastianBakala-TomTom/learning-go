package main

import (
	"fmt"
	"strings"
	"math/rand"
)

// Slices in go are designed to be used by value semantic !!!

func main() {
	// Create slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectFruits(fruits)
	for spoilRandomFruit(fruits) {
		inspectFruits(fruits)
	} 
}

// inspectFruits exposes the slice header for review 
func inspectFruits(slice []string) {
	fmt.Printf("<%p> Length[%d] Capacity[%d]\n", &slice, len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] <%p> %s\n",
			i,
			&slice[i],
			s,
		)
	}
}

func spoilRandomFruit(slice []string) bool {
	canAnythingBeSpoiled := false
	for _, v := range slice {
		if ! strings.Contains(v, "Rotten") {
			canAnythingBeSpoiled = true
			break
		}
	}

	if ! canAnythingBeSpoiled {
		fmt.Println("There is nothing more to spoil.")
		return false
	}

	max := len(slice)
	for true {
		randix := rand.Intn(max)
		fruit := slice[randix]
		if strings.Contains(fruit, "Rotten") { continue }
		slice[randix] = "Rotten " + slice[randix]
		break;
	}

	return true
}