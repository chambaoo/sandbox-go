package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello.")

	user := createUser()

	fmt.Println(user)
	fmt.Printf("from %s: %s\n", user, readMessage(1))
	fmt.Printf("from %s: %s\n", user, readMessage(2))
}

func createUser() string {
	return "Shiba"
}

func readMessage(id int) string {

	messages := [...]string{"Hi.", "Bye."}

	if id < 0 || id > len(messages)-1 {
		// todo: throw error
		return "error: message not found"
	}

	t := time.Now()
	return t.String() + messages[id]
}
