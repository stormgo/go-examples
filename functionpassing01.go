package main

import (
	"fmt"
)

func main() {

	anon := func(name string) string {
		return "Hiya, " + name
	}

	fmt.Println("Hello, playground")
	anotherFunction(anon)
}

func anotherFunction(f func(string) string) {
	result := f("Iris")
	fmt.Println(result) // Prints "Hiya, David"
}
