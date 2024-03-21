package main

import (
	"fmt"
)

func main() {

	var s []string
	fmt.Println("Value s ", s, "Len", len(s))

	s = make([]string, 3)
	fmt.Println("Value s ", s, "Len", len(s))

	for i := 0; i < len(s); i++ {
		s[i] = "a"
	}
	fmt.Println("Value s ", s, "Len", len(s))
	copy(s, c)

}
