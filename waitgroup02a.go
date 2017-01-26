package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://intel.com",
		"http://microsoft.com",
		"http://zrato.com",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		fmt.Println("ok")
		// Launch a goroutine to fetch the URL.
		go process(url)
	}
	// Wait for all HTTP fetches to complete.
	fmt.Println("Ralph")
	wg.Wait()
	fmt.Println("Bye")
}

func process(url string) (ok string, err error) {
	var body string
	// Decrement the counter when the goroutine completes.
	defer wg.Done()
	// Fetch the URL.
	body, err = fetch(url)
	fmt.Println(body)
	return "ok", err
}

func fetch(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Got an error")
		return "", err
	}
	return string(body), nil
}
