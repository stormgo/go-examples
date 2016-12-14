// https://blog.golang.org/go-maps-in-action

package main

import (
	"fmt"
)

var EmptyStruct = struct{}{}

var TargetIds = map[float64]struct{}{}

func init() {
	TargetIds[123.0] = EmptyStruct
	TargetIds[223.0] = EmptyStruct
	TargetIds[333.0] = EmptyStruct
}

func main() {

	ids := map[float64]struct{}{
		8432709.0: EmptyStruct,
		8432758.0: EmptyStruct,
		8432763.0: EmptyStruct,
		8432838.0: EmptyStruct,
		8432857.0: EmptyStruct,
		8432914.0: EmptyStruct,
		8432919.0: EmptyStruct,
	}

	value1, ok1 := ids[8432709.0]
	value2, ok2 := ids[8432710.0]
	fmt.Println(value1, " ", ok1)
	fmt.Println(value2, " ", ok2)

	value3, ok3 := TargetIds[223.0]
	value4, ok4 := TargetIds[224.0]
	fmt.Println(value3, " ", ok3)
	fmt.Println(value4, " ", ok4)
}
