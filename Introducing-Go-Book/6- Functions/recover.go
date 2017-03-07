package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("HELLO THIS IS LIKE PYTHON ASSERT BUT MOST POWERFULL")
}
