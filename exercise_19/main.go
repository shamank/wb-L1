package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {
	var input string

	fmt.Printf("Enter the word: ")
	fmt.Scan(&input)

	runes := []rune(input)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	result := string(runes)
	fmt.Printf("\nReversed of %s is %s", input, result)
}
