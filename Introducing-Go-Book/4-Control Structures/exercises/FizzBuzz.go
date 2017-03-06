package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
			fmt.Println("FizzBuzz")
		}

		if i%3 == 0 {
			fmt.Println(i)
			fmt.Println("Fizz")
		}

		if i%5 == 0 {
			fmt.Println(i)
			fmt.Println("Buzz")
		}

		i += 1
	}
}
