package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	arr := []int{2, 3, 12, 6, 324, 1, 645, 123, 3, 5, -1, 231, 65}

	want := 645

	if res, ok := binarySearch(arr, want); ok {
		fmt.Printf("Index of %d is %d", want, res)
	} else {
		fmt.Printf("Value %d was not found", want)
	}

}

func binarySearch(arr []int, key int) (int, bool) {
	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == key {
			return mid, true
		}
		if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, false
}
