// http://cplus.about.com/od/google-go/a/programming-golang-tutorial-on-functions.htm

package main

import "fmt"

type largearray [1000]int

func Sum(a *largearray) (sum int) {
	for _, value := range *a {
		sum += value
	}
	return
}

func main() {
	var numbers largearray
	for index, _ := range numbers {
		numbers[index] = index
	}

	fmt.Println(Sum(&numbers))
}
