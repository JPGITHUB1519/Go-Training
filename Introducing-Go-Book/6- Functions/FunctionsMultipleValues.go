package main

import (
	"fmt"
)

// return multiple values
func f() (int, int) {
	return 5, 6
}

func main() {
	x, y := f()
	fmt.Println(x, y)
}
