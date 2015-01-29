// http://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/

package main

import "fmt"

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("Woof!")
}

func main() {
	d := &Dog{}
	d.Say()
}
