package main

import "fmt"

/*

   * operator -> access the value of the pointed memory location
   & operator -> access the address of a varibles

*/

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println("Memory Location of X : ", &x)
	fmt.Println("Value of X : ", x)
}
