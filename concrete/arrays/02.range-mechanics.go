package main

import "fmt"


func main() {
	friends := [5]string{"Annie", "Betty", "Charely", "Doug", "Edward"}

	// fmt.Println("+========= Data structure START ===========+")
	// fmt.Println("| index |    address     |   address org   |")
	// fmt.Println("+------------------------------------------+")
	// for i, v := range &friends {
	// 	fmt.Printf("|  [%d]  |  %p  |  %p   |\n", i, &v, &friends[i])
	// }
	// fmt.Println("+=========== Data structure END ===========+\n")
	
	fmt.Printf("Bfr[%s] : ", friends[1])

	// Pointer semantic
	for i := range friends {
		friends[1] = "Jack"
		
		if i == 1 {
			fmt.Printf("Aft[%s]\t\t|+ Pointer Semantic\n", friends[1])
		}

	}

	friends[1] = "Betty"
	fmt.Printf("Bfr[%s] : ", friends[1])

	// Value semantic - will make its own copy of 'friends' variable.
	//   Because of that, any changes in 'friends' var 
	//   will not be reflected to the 'v' var and vice versa 
	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\t\t|+ Value Semantic\n", v)
		}
	}

	friends[1] = "Betty"
	fmt.Printf("Bfr[%s] : ", friends[1])

	// Mixing semantics -- ANTIPATTERN
	//   Using the value semantic form of the for range but with pointer semantic access.
	for i, v := range &friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("Aft[%s]\t\t|+ Mixed Semantic\n", v)

		}
	}

	fmt.Printf("\nAfterAll[%s]\n", friends[1])
}