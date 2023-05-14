package main

import (
	"fmt"
	"unicode"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {

	table := [...]string{"abcd", "abCdefAaf", "aabcd"}

	for _, val := range table {
		fmt.Printf("%s is unique - %v\n", val, checkWord(val))

	}
}

// Сохраняем руны в словарь. Если руна совпадет - return false. Если не было совпадений - return true
func checkWord(word string) bool {
	tmp := make(map[rune]interface{})
	var val rune
	for _, rn := range word {
		val = unicode.ToLower(rn)
		if _, ok := tmp[val]; ok {
			return false
		} else {
			tmp[val] = nil
		}
	}
	return true
}
