// http://www.golang-book.com/8/index.htm

package main

import (
	"fmt"
)

func zeronp(x int) {
	x = 0
}

func zerop(xPtr *int) {
	*xPtr = 0
}


func hundred(xPtr *int) {
	fmt.Println("xPtr = ", xPtr)
	*xPtr = 100
}

func main() {
	x := 5
	zeronp(x)
	fmt.Println(x) // x is 5
	zerop(&x)
	fmt.Println(x) // x is 0

	var p *int
	i := 42
	p = &i
	fmt.Println("p =    ", p)
	hundred(p)
	fmt.Println(i) // i is 100
}
