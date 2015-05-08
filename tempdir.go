package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	datadir, err := ioutil.TempDir("/tmp71", "my-data")
	if err == nil {
		fmt.Println(datadir)
	}
}
