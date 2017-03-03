package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func resta(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(resta(42, 13))
}
