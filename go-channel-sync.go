package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	fmt.Println("Started")
	time.Sleep(time.Second)
	fmt.Println("Ended")

	done <- true
}

func main() {

	
	for i := 0; i < 3; i++ {

		done := make(chan bool, 1)
		
		go worker(done)

	}



	

}
