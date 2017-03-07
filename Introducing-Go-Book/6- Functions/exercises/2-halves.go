package main

import (
	"fmt"
)

func halves(n int) (int, bool) {
	if n%2 == 0 {
		n = n / 2
		return n, true
	} else {
		n = n / 2
		return n, false
	}
}

func main() {
	fmt.Println(halves(1))
}
