package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println("from: ", from, i)
	}
}

func main() {

	f("Direct")

	go f("coroutine")

	go func(msg string) {
		fmt.Print(msg)
	}("hello")

	time.Sleep(time.Second)

	fmt.Println("Done")

}
