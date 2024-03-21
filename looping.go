package main

import "fmt"

func main() {
	fmt.Println("Starting Main go progrma")

	var a [5]int
	fmt.Println("Array value is", a)

	a[1] = 10
	fmt.Println("Array value is", a)
	fmt.Println("Array len", len(a))

	b := []int {1, 2, 3}
	fmt.Println("Array len for b", b)

	var twoD [2][2]int

	for i:= 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i*j
		}
	}

	fmt.Println("Value of twoD: ", twoD)

}
