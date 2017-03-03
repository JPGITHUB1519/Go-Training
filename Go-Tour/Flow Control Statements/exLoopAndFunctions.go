package main

import (
	"fmt"
	"math"
)

func calculateFormula(z, x float64) float64 {
	return z - ((math.Pow(z, 2) - x) / (2 * z))
}

func Sqrt(x float64) float64 {
	res := float64(10)
	for i := 1; i < 5; i++ {
		res = calculateFormula(res, x)
		fmt.Println(res)
	}
	return res
}

func main() {
	//fmt.Println(Sqrt(2))
	//fmt.Println(calculateFormula(10, 612))
	fmt.Println(Sqrt(900))
}

/*
   612 -> 24.738
*/
