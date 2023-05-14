package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала:
//в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	arr := [...]int{2, 4, 6, 8, 10, 12, 14}

	// первый этап - запись в канал из массива
	go func() {
		for _, num := range arr {
			numbers <- num
		}
		close(numbers)
	}()

	// второй этап - запись из первого канала квадрат значения во второй канал
	go func() {
		for num := range numbers {
			squares <- num * num
		}
		close(squares)
	}()

	done := make(chan bool) // для ожидания выполнения горутины

	// третий этап - вывод из второго канала в stdout
	go func() {
		for sq := range squares {
			fmt.Printf("%d ", sq)
		}
		done <- true
	}()
	<-done
}
