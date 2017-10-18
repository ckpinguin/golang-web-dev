package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln(t.doors, "doors opening")
}
func (s sedan) transportationDevice() string {
	return fmt.Sprintln("Luxurious?", s.luxury)
}
func main() {
	t1 := truck{
		vehicle{4, "green"},
		true}
	s1 := sedan{
		vehicle{3, "blue"},
		false}
	fmt.Println(t1.transportationDevice())
	fmt.Println(s1.transportationDevice())
}
