package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// x = g1 // bonks
	x = int(g1) // works
	fmt.Println(x)

}
