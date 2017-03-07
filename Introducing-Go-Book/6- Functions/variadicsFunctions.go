package main

import (
	"fmt"
)

// can take Multiples Arguments
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7))
	// or we can pass slices
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}
