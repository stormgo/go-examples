package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan string) {
	var x string
	for i := 0; ; i++ {
		t := strconv.Itoa(i)
		x = "ping " + t
		c <- x
	}
}
func printer(c chan string) {
	for i := 0; i < 9; i++ {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; i < 3; i++ {
		t := strconv.Itoa(i)
		x := "pong " + t
		c <- x
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	fmt.Println("hola 1")
	go ponger(c)
	fmt.Println("hola 2")
	go printer(c)

	// when the program stops you need to enter something at the command line
	// a carriage return, character or word
	var input string
	fmt.Scanln(&input)

	fmt.Println("hola 3")
}
