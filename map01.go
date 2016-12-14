// https://blog.golang.org/go-maps-in-action

package main

import (
	"fmt"
)

func main() {
	var s struct{}

	ids := map[float64]struct{}{
	    8432709.0: s,
	    8432758.0: s,
	    8432763.0: s,
	}

	value1, ok1 := ids[8432709.0]
	value2, ok2 := ids[8432710.0]
	fmt.Println(value1, " ", ok1)
	fmt.Println(value2, " ", ok2)

}
