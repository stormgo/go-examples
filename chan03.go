package main

import (
	"fmt"
)

func main() {
	ch := numbers()

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

func numbers() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}
