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
func main() {
	x := 5
	zeronp(x)
	fmt.Println(x) // x is 5
	zerop(&x)
	fmt.Println(x) // x is 0
}
