package main

import "fmt"


func main() {
	// Declare a nil slice of strings
	var data []string

	// Capture the capacity of the slice
	lastCap := cap(data)

	// Append about 100k strings to the slice
	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// When the capacity of the slice changes, display the changes
		if lastCap != cap(data) {

			// Calculate the percent of change
			capChg := float64(cap(data) - lastCap) / float64(lastCap) * 100

			lastCap = cap(data)
			
			// Display slice details
			fmt.Printf("Addr[%p]\tIndex[%d]\tCap[%d - %d%%]\n",
				data, record, cap(data), int64(capChg))
		}

	}
}