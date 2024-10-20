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

	fmt.Printf("message sending...\n%s\n", sendMessage("Nagano", ""))
	fmt.Printf("message sending...\n%s\n", sendMessage("", "towel"))
	fmt.Printf("message sending...\n%s\n", sendMessage("Nagano", "towel"))
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
	return t.String() + " " + messages[id]
}

func sendMessage(to string, message string) string {
	if to == "" || message == "" {
		// todo: throw error
		return "error: address or message empty"
	}
	return fmt.Sprintf("to %s: %s", to, message)
}
