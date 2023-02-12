package main

import (
	"fmt"
)

// notifier is an interface that defined notification type behavior.
type notifier interface {
	notify()
}

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
	user  // Embedding
	level string
}

// sendNotification accept values that implement the notifier
// interface and sent notifications
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	// Use of EMBEDDING
	ad := admin{
		user: user{
			name:  "John Smith",
			email: "john.smith@company.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)
}
