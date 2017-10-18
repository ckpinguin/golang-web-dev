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

func main() {
	t1 := truck{vehicle{4, "green"}, true}
	s1 := sedan{vehicle{3, "blue"}, false}
	fmt.Println(t1, s1)
	fmt.Println(t1.color)
	fmt.Println(s1.luxury)
}
