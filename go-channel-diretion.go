package main

import "fmt"

func ping(pings chan string, msg string) {
	pings <- msg
}

func pong(pings chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pingChannel := make(chan string, 1)
	pongChannel := make(chan string, 1)

	go ping(pingChannel, "Hello")
	go pong(pingChannel, pongChannel)

	fmt.Println(<-pongChannel)

}
