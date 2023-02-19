// This is an example of using composition and interfaces.
// This is something we want to do in Go.
// We will group types by their behavior and not by their state.
// This pattern does provide a good design pattern principle in a Go program.

package main

import "fmt"

// Speaker provide a common behavior for all concrete types to follow
// if they wan to be a part of this group.
// This is a contract for these concrete types to follow.
type Speaker interface {
	Speak()
}

// Dog contains everytihg a Dog needs.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d *Dog) Speak() {
	fmt.Printf(
		"Woof! My name is %s, it is %t I am a mammal with a pack factor of %d\n",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

// Cat contains everyting a Cat needs.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (d *Cat) Speak() {
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d\n",
		d.Name,
		d.IsMammal,
		d.ClimbFactor,
	)
}

func main() {
	speakers := []Speaker{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		&Dog{
			Name:       "Crudo",
			IsMammal:   true,
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes
		&Cat{
			Name:        "Prosciutto",
			IsMammal:    true,
			ClimbFactor: 12,
		},
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
