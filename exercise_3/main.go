package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	// waitGroup + Mutex
	var (
		mx  sync.Mutex
		wg  sync.WaitGroup
		res int
	)

	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			mx.Lock()
			res += num * num
			mx.Unlock()
			wg.Done()
		}(num)
	}
	wg.Wait()

	fmt.Printf("result with mutex: %d", res)

	fmt.Print("\n-------------\n")

	// with Channel
	sq := make(chan int)

	for _, num := range arr {
		go func(num int) {
			sq <- num * num
		}(num)
	}

	res = 0

	for _ = range arr {
		res += <-sq
	}
	fmt.Printf("result with chan: %d", res)

	fmt.Print("\n-------------\n")

	// with atomic pkg + channel
	var res_int64 int64

	ch := make(chan interface{})

	for _, num := range arr {
		go func(num int64) {
			atomic.AddInt64(&res_int64, num*num)
			ch <- nil
		}(int64(num))
	}
	for _ = range arr {
		<-ch
	}

	fmt.Printf("result with atomic pkg: %d", res_int64)

}
