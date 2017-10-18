package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type secretagent struct {
	person person
	status string
	no     int
}

func (s secretagent) saSpeak() {
	fmt.Println("My name is...", s.person.name, "...James ", s.person.name)
}

func (p person) pSpeak() {
	fmt.Println("My name is ", p.name)
}

func main() {
	p1 := person{name: "Hubi", age: 42}
	s1 := secretagent{person: p1, status: "active", no: 7}
	fmt.Println(p1.name)
	p1.pSpeak()
	fmt.Println(s1.no)
	s1.saSpeak()
	s1.person.pSpeak()
}
