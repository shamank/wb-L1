package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {

	// ОСТАНОВКА ГОРУТИНЫ С ПОМОЩЬЮ КАНАЛА
	done := make(chan bool)
	go func(done chan bool) {
		fmt.Println("Goroutine with channel (1) started...")
		for {
			select {
			case <-done:
				fmt.Println("Goroutine with channel was stopped..")
				return
			default:
				// что-то делаем..
			}
		}

	}(done)
	// запишем в канал информацию о том, что мы закончили
	go func() {
		time.Sleep(time.Second)
		done <- true
	}()
	time.Sleep(time.Second * 2)

	// ОСТАНОВКА ГОРУТИНЫ С ПОМОЩЬЮ КАНАЛА (VERSION 2)
	ch := make(chan interface{})

	go func(ch chan interface{}) {
		fmt.Println("Goroutine with channel (2) started...")
		for _ = range ch {
			// что-то делаем полезное..

			// цикл остановится только при закрытии канала
			// если канал не закрыть, и никто не будет писать - то будет deadlock
		}
		fmt.Println("Goroutine with channel #2 was stopped..")
	}(ch)
	// закрываем канал (передать какое-то значение надо, чтобы он сразу не deadlock'нулся)
	go func() {
		time.Sleep(time.Second)
		ch <- true
		close(ch)

	}()
	time.Sleep(time.Second * 2)

	// ОСТАНОВКА ГОРУТИНЫ С ПОМОЩЬЮ CONTEXT
	ctx, cancel := context.WithCancel(context.Background()) // еще можно использовать контекст с timeout...
	go func(ctx context.Context) {
		fmt.Println("Goroutine with context started..")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine with context was stopped..")
				return
			default:
				// снова что-то очень полезное делаем..
			}
		}
	}(ctx)
	// остановим контекст
	go func(cancel context.CancelFunc) {
		time.Sleep(time.Second)
		cancel()
	}(cancel)
	time.Sleep(time.Second * 2)
}
