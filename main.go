package main

import (
	"fmt"
	"reflect"
	)

func main() {

	var salary float32 = 1000.981111;
	fmt.Printf("Name: %v", reflect.TypeOf(salary));
	fmt.Println()
	
	var age int = 30;
	fmt.Printf("Age is %d", age);

	if age >= 60 {
		fmt.Println("Senior citizen")	
	} else {
		fmt.Println("Citizen")
	}
	
}
