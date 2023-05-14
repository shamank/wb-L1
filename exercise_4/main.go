package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

func main() {
	var (
		wg           sync.WaitGroup
		workersCount int
	)
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(500 * time.Millisecond)

	data := make(chan interface{}) // канал с произвольными данными

	fmt.Printf("Enter the number of workers: ")
	if _, err := fmt.Scan(&workersCount); err != nil {
		log.Fatal(err.Error())
	}

	// запуск воркеров
	for i := 0; i < workersCount; i++ {
		go newWorker(ctx, &wg, i, data)
	}

	// создание канала, в который запишется сигнал из ОС
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// для проверки работы воркеров отправляем в канал рандомные числа
	// если пришел сигнал о завершнии работы, то останавливаем воркеры
	for {
		select {
		case <-quit:
			stopWorkers(cancel, &wg)
			close(data)
			return
		case <-ticker.C:
			data <- rand.Float64()
		}
	}
}

func newWorker(ctx context.Context, wg *sync.WaitGroup, workerNum int, data chan interface{}) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // завершение работы воркера
			fmt.Printf("\n[Goroutine %d]: End...\n", workerNum)
			return
		case d := <-data:
			fmt.Printf("\n[Goroutine %d]: %v\n", workerNum, d)
		}
	}
}

func stopWorkers(cancel context.CancelFunc, wg *sync.WaitGroup) {
	fmt.Printf("\nreceived a signal to terminate the program\n")
	cancel()
	wg.Wait()
	fmt.Printf("program completed successfully!")
}
