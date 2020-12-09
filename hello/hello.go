package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Rbaodf")
	fmt.Println(message)
}
