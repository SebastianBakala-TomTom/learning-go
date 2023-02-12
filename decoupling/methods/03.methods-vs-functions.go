package main

import (
	"fmt"
)

type data struct {
	name string
	age  int
}
// 'bill' type cannot be used in any method which is using the 'data' type
type bill data

// displayName provides a pretty print view of the name
func (d data) displayName() {
	fmt.Println("My name is", d.name)
}

// setAge sets the age and displays the value
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println("My age is", d.age)
}

func main() {
	b := data{
		name: "Bill",
	}

	fmt.Println("Proper calls to methods!")

	// How we actuall call methods in Go.
	b.displayName()
	b.setAge(24)

	fmt.Println("\nThis is what Compiler is Doing:")
	// This is what Go is doing underneath.
	// NEVER CALL METHODS LIKE THAT. ITS ONLY TO SHOW WHAT IS GOING ON
	data.displayName(b)
	(*data).setAge(&b, 24)
	
	// functions in go are values
	// we can do this!
	// f1 is a reference type
	// because displayName reciever is value type
	// f1 points to its own copy of b value
	f1 := b.displayName

	// call the method via the variable
	f1()
	b.name = "Drake"

	// becase setAge reciever is pointer type
	// f2 points to the original data!!!
	f2 := b.setAge
	f2(52)
	f1()

}
