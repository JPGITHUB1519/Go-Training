package main

import "fmt"

func main() {
	var fah float64
	var cel float64
	fmt.Println("Type the Fahrenheit Grades : ")
	fmt.Scanf("%f", &fah)
	cel = (fah - 32) * (float64(5) / float64(9))
	fmt.Printf("%f\n", cel)
}
