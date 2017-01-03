package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	results := mystrings()

	for result := range results {
		fmt.Println(strings.Join(result, " "))
	}
}

func getString(i int) (strary []string) {
	var buffer bytes.Buffer
	start := "john ran home "
	buffer.WriteString(start)
	buffer.WriteString(strconv.Itoa(i))
	end := buffer.String()
	strary = strings.Split(end, " ")
	return strary
}

func mystrings() <-chan []string {
	ch := make(chan []string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- getString(i)
		}
		close(ch)
	}()

	return ch
}
