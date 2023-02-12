package main

import (
	"fmt"
)


type user struct {
	name	string
	email	string
}

// nofify() implements a method with a value reciever 
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email,
	)
}

// changeEmail() implements a method with a pointer reciever
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// Values of type user, can be used to call methods
	// declared with both value and pointer recievers
	bill := user{"Bill", "bill@gmail.com"}
	// go will convert this value to the pointer because 
	// changeEmail method wants a reciever to be a pointer
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer recievers
	joan := &user{"Joan", "joan@gmail.com"}
	joan.changeEmail("joan@hotmail.com")
	// go will convert this pointer to the value because
	// notify method wants a reciever to be a value
	// WARNING! Not every data can be converted this way
	// 	According to this, we cannot be 100% sure 
	//  that we are operating on the copy rather than the original one 
	joan.notify()

	users := []user{
		{"Ed", "ed@gmail.com"},
		{"Erick", "erick@gmail.com"},
	}

	fmt.Println("*************************")
	// this is bad!
	// we are using copy of users and modyfing copies only!
	// original data are not touched!
	for _, v := range users {
		domain := "hotmail.com"
		v.changeEmail(v.name + "@" + domain)
	}
	// proof
	for i := range users {
		users[i].notify()
	}

	fmt.Println("*************************")
	// this is the reason why choosing the right semantics is so important
	// we should not copy users but operate on original data instead
	for i := range users {
		domain := "hotmail.com"
		users[i].changeEmail(users[i].name + "@" + domain)
	}
	// proof
	for i := range users {
		users[i].notify()
	}
}