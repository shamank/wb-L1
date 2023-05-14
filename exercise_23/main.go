package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	i := 3
	fmt.Printf("[Input] %v\n", slice)

	slice = append(slice[:i], slice[i+1:]...)
	fmt.Printf("[Delete with append] %v\n", slice)

	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	fmt.Printf("[Delete with copy] %v\n", slice)

	// здесь меняется порядок, поэтому может быть неактуально
	slice[i] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	fmt.Printf("[Delete with move to last] %v\n", slice)
}
