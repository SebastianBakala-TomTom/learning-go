package main

import (
	"fmt"

	"github.com/SebastianBakala-TomTom/learning-go/decoupling/exporting/example_5/users"
)

func main() {

	// Create a value of type Manager from the users package.
	user := users.Manager{
		Title: "Dev Manager",
	}

	// Set the exported fields from the unexported user inner type.
	user.Name = "Alice"
	user.ID = "13"

	fmt.Printf("User: %#v\n", user)
}
