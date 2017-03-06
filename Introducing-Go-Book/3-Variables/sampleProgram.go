package main

import "fmt"

func main() {
	fmt.Print("Enter a Number : ")
	var input float64
	// read from keyboard
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
