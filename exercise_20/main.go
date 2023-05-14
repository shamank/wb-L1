package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	var input string

	//fmt.Printf("Enter the sentence: ")
	//fmt.Scan(&input)

	input = "sun dog snow"

	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	fmt.Printf("%v", words)

}
