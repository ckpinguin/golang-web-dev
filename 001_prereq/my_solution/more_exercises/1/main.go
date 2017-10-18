package main

import (
	"fmt"
)

func main() {
	s1 := []int{0, 1, 2, 3, 12}
	fmt.Println(s1)

	for k, _ := range s1 {
		fmt.Println(k)
	}

	for k, v := range s1 {
		fmt.Println(k, "=", v)
	}
}
