// http://stackoverflow.com/questions/25343853/golang-method-receivers

package main

import "fmt"
import "math"

type MyFloat float64

func (x MyFloat) Abs() float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
