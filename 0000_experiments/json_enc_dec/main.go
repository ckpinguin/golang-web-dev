package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	testFun("World")
	f, err := os.Create("test.json")
	check(err)
	defer f.Close()
	json.NewEncoder(f).Encode(testFun)
}

func testFun(s string) string {
	log.Println("Hello", s)
	return s + s
}

func check(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}
