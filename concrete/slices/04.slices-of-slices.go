package main

import "fmt"

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

	// create a sub-slice
	fruits2 := fruits[2:4]

	inspectFruits(fruits2)

	// Both slices are using the same addresses for the elements that they coresponds to
	// SIDE EFFECT - overriding some value in one, will be reflected in the other (if overlaps)
	fruits2[0] = "TAKEN AWAY"

	inspectFruits(fruits)
	inspectFruits(fruits2)


	// To avoid it, we can use go mechanisms
	// PROBLEM: When we not define custom capacity, go will use the capacity from the original slice, 
	// but it will start from the starting index that we are using to create a sub-slice
	// SOLUTION: Go enables to define sub-slices with custom capacity
	fruits3 := fruits[3:5:5]

	// it have the same address space
	inspectFruits(fruits3)

	// but we can append to this slice and go will allocate new backing array
	fruits3 = append(fruits3, "Mango")
	inspectFruits(fruits3)

	// another option is to use go built-in 'copy' fucntion
	fruits4 := make([]string, 4, 8)
	copy(fruits4, fruits2)

	inspectFruits(fruits4)
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