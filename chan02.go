package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	fmt.Println("hola 1")
	go ponger(c)
	fmt.Println("hola 2")
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
