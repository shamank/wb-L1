package main

import "fmt"

//Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

func main() {
	var (
		number int64
		index  int64
	)
	fmt.Print("Enter the number: ")
	fmt.Scan(&number)

	fmt.Print("\nEnter the index of Bit: ")
	fmt.Scan(&index)
	index--

	fmt.Printf("\nGet number: %d [%b]\n", number, number)

	// побитовое (ИЛИ)
	number = number | (1 << index)
	fmt.Printf("\nChange %d bit on 1: %d [%b]", index, number, number)

	// сброс бита (И НЕ)
	number = number &^ (1 << index)
	fmt.Printf("\nChange %d bit on 0: %d [%b]", index, number, number)
}
