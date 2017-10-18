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

type human interface {
	speak()
}

func (s secretagent) speak() {
	fmt.Println("My name is...", s.person.name, "...James ", s.person.name)
}

func (p person) speak() {
	fmt.Println("My name is ", p.name)
}

func hSpeak(h human) {
	h.speak()
}

func main() {
	p1 := person{name: "Hubi", age: 42}
	s1 := secretagent{person: p1, status: "active", no: 7}
	hSpeak(p1)
	hSpeak(s1)
	// fmt.Println(p1.name)
	// p1.pSpeak()
	// fmt.Println(s1.no)
	// s1.saSpeak()
	// s1.person.pSpeak()
}
