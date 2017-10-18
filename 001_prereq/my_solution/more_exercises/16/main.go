package main

import "fmt"

type gator int
type flamingo bool
type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	bayou(g1)

	var f1 flamingo
	fmt.Println(f1)
	fmt.Printf("%T\n", f1)
	bayou(f1)
}
