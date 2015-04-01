// http://cplus.about.com/od/google-go/a/Learn-To-Program-Go-Tutorial-Four.htm

package main

import "fmt"

func main() {

	// pointer to a list of five ints
	var p2list *[5]int

	// a list of five pointers to ints
	var listofp [5]*int

	p2list = new([5]int)
	for i := 0; i < 5; i++ {
		listofp[i] = new(int)
		*listofp[i] = i * i * i
		(*p2list)[i] = i * i
	}

	fmt.Println(p2list)
	fmt.Println(listofp)
	for i := 0; i < 5; i++ {
		fmt.Print(*listofp[i], " ")
	}
	fmt.Println("")
}
