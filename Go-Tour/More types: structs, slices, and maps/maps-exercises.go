package main

import (
	"fmt"
	"strings"
)

func WordRepetition(search string, line string) int {
	// return the cant of apparitions of a word in a line
	// search -> word to count, line -> the line to lookup in
	arr := strings.Split(line, " ")
	countApparitions := 0
	for i := 0; i < len(arr); i++ {
		if search == arr[i] {
			countApparitions++
		}
	}
	return countApparitions

}

// word : repetition_of_that_world
func WordCount(s string) map[string]int {
	dic := make(map[string]int)
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); i++ {
		dic[arr[i]] = WordRepetition(arr[i], s)
	}
	return dic
}

func main() {
	// WordCount := "Hola Mi Nombre es Jean"
	// wc.Test(WordCount)
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
