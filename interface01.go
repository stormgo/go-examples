package main

import "fmt"

type dog struct {
	name     string
	isabitch bool
}

func (d dog) bark() {
	fmt.Println(d.name, "Goes: 'Woof, Woof!'")
}

func (d dog) move() {
	fmt.Println(d.name, "Moves noisily")
}

type cat struct {
	name string
	age  int
}

func (c cat) move() {
	fmt.Println(c.name, "Walks very delicately")
}

type doggy interface {
	bark()
}

type mammal interface {
	move()
}

func main() {
	fido := dog{"Fido", false}
	max := cat{"Max", 5}
	mammal1 := mammal(fido)
	mammal1.move()
	fido.bark()
	mammal2 := mammal(max)
	mammal2.move()
	fmt.Println(mammal1)
	fmt.Println(mammal2)
}
