package main

import "fmt"

func main() {
	// to be run after the wrapper function complete
	defer fmt.Println("world")
	fmt.Println("hello")
}
