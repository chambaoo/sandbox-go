package lib

import (
	"fmt"
	"time"
)

func ReadMessage(id int) string {

	messages := [...]string{"Hi.", "Bye."}

	if id < 0 || id > len(messages)-1 {
		// todo: throw error
		return "error: message not found"
	}

	t := time.Now()
	return t.String() + " " + messages[id]
}

func SendMessage(to string, message string) string {
	if to == "" || message == "" {
		// todo: throw error
		return "error: address or message empty"
	}
	return fmt.Sprintf("to %s: %s", to, message)
}
