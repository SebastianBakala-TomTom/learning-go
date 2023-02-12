package main

import "fmt"


func printFruits() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Value semantic
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

func printNumbers() {
	numbers := [4]int{20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

func main() {
	printFruits()
	printNumbers()
}