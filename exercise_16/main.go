package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	arr := []int{1, 2, 33, 23, 342, 2343, 2, 2, 3, 5, 3, 6, 2132, 7243}

	fmt.Printf("Before sort: %v\n", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Printf("After sort: %v\n", arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// разбиение ламуто
func partition(arr []int, low, high int) int {
	p := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
