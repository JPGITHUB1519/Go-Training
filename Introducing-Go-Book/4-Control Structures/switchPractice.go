package main

import "fmt"

func main() {

	var op int

	fmt.Printf("Opciones \n\n")
	fmt.Printf("1- Sumar\n")
	fmt.Printf("2- Restar\n")
	fmt.Printf("3- Multiplicar\n")
	fmt.Printf("4- Dividir\n\n")

	fmt.Printf("Elija una Opcion : ")
	fmt.Scanf("%d", &op)
	fmt.Println(op)
	switch op {
	case 1:
		fmt.Println("Haz Elegido Sumar")
	case 2:
		fmt.Println("Haz Elegido Restar")
	case 3:
		fmt.Println("Haz Elegido Multiplicar")
	case 4:
		fmt.Println("Haz Elegido Dividir")
	default:
		fmt.Println("Unknow Number")
	}
}
