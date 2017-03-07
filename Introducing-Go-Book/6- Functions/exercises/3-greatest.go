package main

import "fmt"

func greatest(args ...int) int {
	if len(args) < 0 {
		panic("You Must Pass a Parameter")
	}
	max := args[0]
	if len(args) > 1 {
		for i := 1; i <= len(args); i++ {
			if i > max {
				max = i
			}
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(1))
}
