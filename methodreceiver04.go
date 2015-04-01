// http://cplus.about.com/od/google-go/a/programming-golang-tutorial-functions-methods.htm

package main

import "fmt"

type dog struct {
	name     string
	isabitch bool
}

type showdog struct {
	adog      dog
	groomedby string
}

func (d dog) bark() {
	fmt.Println(d.name, "Goes: 'Woof, Woof!'")
}

func (d showdog) bark() {
	fmt.Println(
		d.adog.name,
		"(Groomed by ",
		d.groomedby,
		") Goes: 'Woof, Woof!'")
}

func (d dog) playdead() {
	fmt.Print(d.name, " rolls over, lies on ")
	if d.isabitch {
		fmt.Print("her back")
	} else {
		fmt.Print("his back")
	}
	fmt.Println(", with legs in the air... awwww!!")
}

func (d *dog) rename(newname string) {
	fmt.Println((*d).name, " will forever more answer to", newname)
	(*d).name = newname
}

func main() {
	fido := showdog{dog{"Fido", false}, "Alphonse"}
	fido.bark()
	fido.adog.playdead()
	lulu := dog{"Lulu", true}
	lulu.playdead()
	lulu.rename("Lola")
	lulu.bark()
}
