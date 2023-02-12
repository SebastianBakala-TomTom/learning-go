package main

import (
	"fmt"
	"sort"
)

type user struct {
	name 	string
	surname string
}

func main() {
	// Create a map of users with a key of type string
	users := make(map[string]user)

	// Add key/value pairs to the map
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Peter", "Jackson"}

	for key, value := range users {
		fmt.Println(key, value)
	}
	fmt.Println()

	for key := range users {
		fmt.Println(key)
	}
	fmt.Println()

	for _, value := range users {
		fmt.Println(value.name)
	}
	fmt.Println()

	// delete user
	delete(users, "Roy")

	// find user
	u, found := users["Roy"]

	fmt.Println("Roy", found, u)
	fmt.Println()
	fmt.Println()

	// Sort map for predictable order
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}
	// Sort keys alphabetically
	sort.Strings(keys)

	// Output map using keys as an iterator
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}