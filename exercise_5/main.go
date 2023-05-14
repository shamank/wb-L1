package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {

	ch := make(chan interface{})
	done := make(chan bool)

	var N time.Duration = 1
	timeout := time.After(N * time.Second)

	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("closing goroutine on write\n")
				return
			default:
				ch <- rand.Float64()

			}
		}
	}()

	go func() {
		for c := range ch {
			fmt.Printf("new message: %v\n", c)
		}
		fmt.Printf("closing goroutine on read\n")
		return
	}()

	for {
		select {
		case <-timeout:
			fmt.Printf("Timeout!!!\n")
			done <- true // сообщаем первой горутине (чтобы больше не записывала)
			close(ch)    // сообщаем второй горутине (что больше не запишется ничего)
			fmt.Printf("program completed successfully!\n")
			return
		default:
		}
	}

}
