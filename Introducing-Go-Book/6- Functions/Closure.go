package main

import "fmt"

// functions inside functions
func main() {
	// asigning variable to local functions
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}
