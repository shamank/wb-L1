package main

import (
	"fmt"
	"math"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.

func main() {
	temps := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := make(map[int][]float64)
	inputTemps := make(chan float64)

	go func() {
		for _, tmp := range temps {
			inputTemps <- tmp
		}
		close(inputTemps)
	}()

	done := make(chan bool)
	go func() {
		var key int
		for tmp := range inputTemps {
			if tmp >= 0 {
				key = int(math.Floor(tmp/10)) * 10
			} else {
				key = int(math.Ceil(tmp/10)) * 10
			}
			result[key] = append(result[key], tmp)
		}
		done <- true
	}()
	<-done
	fmt.Printf("%v\n", result)

	// либо просто в цикле....
	result_2 := make(map[int][]float64)
	var key int
	for _, tmp := range temps {
		if tmp >= 0 {
			key = int(math.Floor(tmp/10)) * 10
		} else {
			key = int(math.Ceil(tmp/10)) * 10
		}
		result_2[key] = append(result_2[key], tmp)
	}
	fmt.Printf("%v\n", result_2)
}
