package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	g1.greeting()
}
