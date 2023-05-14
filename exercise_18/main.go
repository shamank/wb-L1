package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	value int64
	mx    *sync.Mutex
}

func main() {
	var mx sync.Mutex
	cnt := &Counter{value: 0, mx: &mx} // с использованием Mutex
	cnt2 := &Counter{}                 // с использованием пакета Atomic

	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {

			defer wg.Done()

			cnt.Inc(1)

			atomic.AddInt64(&cnt2.value, 1)

		}(&wg)
	}
	wg.Wait()
	fmt.Printf("%d\n", cnt.value)
	fmt.Printf("%d\n", cnt2.value)

}

func (c *Counter) Inc(delta int64) {
	c.mx.Lock()
	c.value += delta
	c.mx.Unlock()
}
