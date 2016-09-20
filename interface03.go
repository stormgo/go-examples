package main

import "fmt"

type Walker interface {
	Walk(miles int)
}

type Camel struct {
	Name string
}

func (c Camel) Walk(miles int) {
	fmt.Printf("%s is walking %v miles\n", c.Name, miles)
}

func LongWalk(w Walker) {
	w.Walk(500)
	w.Walk(500)
}

func main() {
	c := Camel{"Bill"}
	LongWalk(c)
}
