package main

import "fmt"


type user struct {
	likes int
}

func main() {
	users := make([]user, 3)

	sharedUser := &users[1]

	sharedUser.likes++

	for i := range users {
		fmt.Printf("User: %d\tLikes: %d\n", i, users[i].likes)
	}

	fmt.Println("\n# Side effects below")
	// now we will create a side effect which leads to memory leak

	// appending the slice when its length equals to the capacity will create new data structure
	users = append(users, user{})

	// new data structure means the new address space
	// but our sharedUser will still point to the old address
	sharedUser.likes++

	// this leads to the situation when garbage collector will not clean the old backing array ever
	// and appending likes to the sharedUser will not mutate the new data structure
	for i := range users {
		fmt.Printf("User: %d\tLikes: %d\n", i, users[i].likes)
	}

	fmt.Println("Be carefull. Although I will fall into this trap in the future anyways :D.")
}
