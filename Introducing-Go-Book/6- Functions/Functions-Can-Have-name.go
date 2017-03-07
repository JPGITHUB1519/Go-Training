package main

import "fmt"

// we can put name to the return value
func f2() (r int) {
	r = 1
	return
}

func main() {
	fmt.Println(f2())
}
