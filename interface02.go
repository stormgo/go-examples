// http://cplus.about.com/od/google-go/a/google-go-tutorial-nine-interfaces.htm

package main

import "fmt"

type dog struct {
	name     string
	isabitch bool
}

func (d dog) bark() {
	fmt.Println(d.name, "Goes: 'Woof, Woof!'")
}

type puppy struct {
	name string
	age  int
}

func (p puppy) bark() {
	fmt.Println(p.name, "Goes: 'Squeak, Squeak!'")
}

type doggy interface {
	bark()
}

func main() {
	fido := dog{"Fido", false}
	pip := puppy{"Pip", 5}
	pack := [...]doggy{fido, pip}
	for _, d := range pack {
		fmt.Println(doggy(d))
		d.bark()
	}
}
