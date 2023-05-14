package main

import (
	"sync"
)

// Реализовать конкурентную запись данных в map.
func main() {
	var (
		mx     sync.Mutex
		mx_map sync.Map
	)
	data := make(map[int]interface{})

	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {

		// можно использовать Mutex для безопасной записи в мапу
		go func(i int, wg *sync.WaitGroup) {
			wg.Add(1)
			set(i, struct{}{}, data, &mx)
			wg.Done()
		}(i, &wg)

		// либо же - использовать готовый тип sync.Map
		go func(i int, wg *sync.WaitGroup) {
			wg.Add(1)
			mx_map.Store(i, struct{}{})
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()

}

func set(key int, value interface{}, data map[int]interface{}, mx *sync.Mutex) {
	mx.Lock()
	data[key] = value
	mx.Unlock()
}
