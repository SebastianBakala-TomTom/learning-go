package main

import (
	"fmt"

	"github.com/SebastianBakala-TomTom/learning-go/decoupling/exporting/example_4/users"
)

func main() {
	user := users.User{
		Name:     "Alice",
		ID:       "13",
		password: "*******",
	}

	// NOT EXPORTED FIELD
	// ./example.go:16: unknown users.User field 'password' in structure User

	fmt.Printf("User: %#v\n", user)
}
