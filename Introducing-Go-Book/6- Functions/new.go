package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	// new allocates enough memory to fit a value of that type, and returns a pointer to it.
	// xPtr is only a pointer referring to a memory value() although in go is 0 by default
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
