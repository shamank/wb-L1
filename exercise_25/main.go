package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	fmt.Println("Start sleep by chan..")
	sleepByChan(2 * time.Second)
	fmt.Println("WakeUp after sleep")

	fmt.Println("Start sleep by cycle..")
	sleepByCycle(2 * time.Second)
	fmt.Println("WakeUp after sleep")
}

// можно также через context WithTimeout
func sleepByChan(t time.Duration) {
	timeout := time.After(t)
	<-timeout
	return
}

func sleepByCycle(t time.Duration) {
	s := time.Now()
	for {
		if time.Since(s) >= t {
			return
		}
	}
}
