package main

import (
	"fmt"
	"mokume/lib"
)

func main() {
	fmt.Println("Hello.")

	user := createUser()

	fmt.Println(user)
	
	fmt.Printf("from %s: %s\n", user, lib.ReadMessage(1))
	fmt.Printf("from %s: %s\n", user, lib.ReadMessage(2))

	fmt.Printf("message sending...\n%s\n", lib.SendMessage("Nagano", ""))
	fmt.Printf("message sending...\n%s\n", lib.SendMessage("", "towel"))
	fmt.Printf("message sending...\n%s\n", lib.SendMessage("Nagano", "towel"))
}

func createUser() string {
	return "Shiba"
}
