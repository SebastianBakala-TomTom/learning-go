package main

import "fmt"


func main() {
	// Create slice with a length of 5 elements.
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Try to add a value above the slice length.
	fruits[5] = "Spoiled Mango"
	// panic: runtime error: index out of range

	fmt.Println(fruits)
}