package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {

		for {
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Jobs completed")
				done <- true
				return
			}
		}

	}()

	for i := 0; i < 10; i++ {
		jobs <- i
		fmt.Println("Sent Job", i)
	}

	close(jobs)
	fmt.Println("Done sending jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
