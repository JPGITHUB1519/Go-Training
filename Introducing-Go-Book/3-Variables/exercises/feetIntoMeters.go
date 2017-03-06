package main

import "fmt"

func main() {
	// 1 ft = 0.3048 m
	var feets float64
	fmt.Print("Enter Feets : ")
	fmt.Scanf("%f", &feets)
	meters := feets * 0.3048
	fmt.Printf("Meters : %f\n", meters)
}
