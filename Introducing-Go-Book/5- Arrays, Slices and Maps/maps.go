package main

import "fmt"

// maps are like dictionary

func main() {
	// dic {"string" : int}
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}
