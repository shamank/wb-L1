package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	arr := [...]string{"cat", "cat", "dog", "cat", "tree"}
	result := make(map[string]interface{})

	for _, str := range arr {
		result[str] = nil
	}

	sliceResult := []string{}
	for str := range result {
		sliceResult = append(sliceResult, str)
	}

	fmt.Printf("%v", sliceResult)
}
