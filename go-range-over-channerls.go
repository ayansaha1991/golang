package main

import "fmt"

func main() {

	queue :=  make(chan string)
	queue <- "one"
	queue <- "second"
	close(queue)

	for v := range queue {
		fmt.Println(v)
	}


}