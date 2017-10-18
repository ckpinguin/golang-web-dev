package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	p1 := person{"Chrigi", "Kälin"}
	fmt.Println(p1)
	fmt.Println(p1.fName)
}
