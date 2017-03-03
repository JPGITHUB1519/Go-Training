package main

import "fmt"

func main() {
	// variables declared without a value are given their to zero
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}
