package main

import (
	"fmt"
)

// struct defining a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method notifies users of different events
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email,
	)
}

type admin struct {
	person user // NOT Embedding
	level  string
}

type adminEmbedded struct {
	user  // Embedding
	level string
}

func main() {
	// Create an admin user
	ad := admin{
		person: user{
			name:  "John Smith",
			email: "john.smith@company.com",
		},
		level: "super",
	}

	// We can access field methods
	ad.person.notify()

	// Use of EMBEDDING
	ade := adminEmbedded{
		user: user{
			name:  "John Smith",
			email: "john.smith@company.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ade.user.notify()

	// The inner type's method is promoted -- Because of embedding
	ade.notify()
}
