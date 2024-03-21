package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	go func ()  {
		messages <- "something for you"
	}()

	message := <- messages
	fmt.Println("Received from channel: ", message)
}
