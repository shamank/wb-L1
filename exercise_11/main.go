package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}
	result := intersects(a, b)
	fmt.Printf("%v", result)

}

func intersects(a, b []int) []int {
	set := make(map[int]bool)
	res := []int{}
	for _, num := range a {
		set[num] = true
	}

	for _, num := range b {
		if set[num] {
			res = append(res, num)
		}
	}
	return res
}
