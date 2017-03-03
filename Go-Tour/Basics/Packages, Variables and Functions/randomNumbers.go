package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand sedd fedding
	rand.Seed(time.Now().UTC().UnixNano())
	for a := 0; a < 10; a++ {
		fmt.Println(rand.Intn(10))
	}
}
