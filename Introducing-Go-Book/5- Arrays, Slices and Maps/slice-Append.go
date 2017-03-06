package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(slice1, slice2)
}
