// http://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/

package main

import "fmt"

type Walker1 interface {
	Walk1(miles int)
}

type Walker2 interface {
	Walk2(miles int)
}

type Camel struct {
	Name string
}

func (c Camel) Walk1(miles int) {
	fmt.Printf("%s is walking %v miles\n", c.Name, miles)
}

func (c *Camel) Walk2(miles int) {
	fmt.Printf("%s is walking %v miles\n", c.Name, miles)
}

func LongWalk1(w Walker1) {
	w.Walk1(500)
}

func LongWalk2(w Walker2) {
	w.Walk2(600)
}

func main() {
	c1 := Camel{"Bill"}
	c2 := &Camel{"Jane"}
	LongWalk1(c1)
    LongWalk2(c2)
}
