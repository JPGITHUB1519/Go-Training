package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	// move the call to second to the end of the functions
	// it is most use when closing file
	// runs even if a panic ocurs
	defer second()
	first()
}
