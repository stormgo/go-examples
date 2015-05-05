package main

import (
  "fmt"
	"net/http"
  "sync"
)

// main is the starting point of the program
func main() {

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://btcsweet.com/",
    //"http://localhost:8080",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			// defer wg.Done()
			// Fetch the URL.
			stuff, err := http.Get(url)
      if err == nil {
        fmt.Println(url,"\n")
        fmt.Println(stuff)
        fmt.Println("\n")
        wg.Done()
      }
		}(url)
	}
	// Wait for all HTTP fetches to complete.
  fmt.Println("waiting for everyone to finish")
	wg.Wait()
}
