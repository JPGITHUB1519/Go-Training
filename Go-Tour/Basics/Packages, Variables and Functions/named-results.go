package main

import "fmt"

// we can put name to the return value
func split(sum int) (x, y int) {
	// variables must be in the top for naked returns
	x = sum * 4 / 9
	y = sum - x
	// naked return
	return
}

func main() {
	fmt.Println(split(17))
}
