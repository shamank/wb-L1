package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	//Использование sync.WaitGroup
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			pprint(n)
		}(num)
	}
	wg.Wait()

	fmt.Print("---------------------\n")

	// Использование каналов
	ch := make(chan int)

	for _, num := range arr {
		go func(n int) {
			ch <- n
		}(num)
	}

	for i := 0; i < len(arr); i++ {
		pprint(<-ch)
	}

}

func pprint(n int) {
	fmt.Printf("%d^2 = %d\n", n, n*n)
}
