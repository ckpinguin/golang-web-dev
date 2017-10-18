package main

import "fmt"

func main() {
	m1 := map[string]int{"dog": 0, "cat": 42}
	fmt.Println(m1)
	for k := range m1 {
		fmt.Println(k)
	}
	for k, v := range m1 {
		fmt.Println(k, "=", v)
	}
}
