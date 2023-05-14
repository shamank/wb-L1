package main

import (
	"fmt"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// решение первой проблемы
func yetSomeFunc() {
	v := createHugeString(1 << 10)
	res := []rune(v)
	justString = string(res[:100])
}

// решение еще одной проблемы, если нас интересует не число символов, а байты
func yetYetSomeFunc() {
	v := createHugeString(1 << 10)
	res := make([]byte, 100)
	copy(res, v)
	justString = string(res)
}

func main() {

	someFunc()
	fmt.Printf("result of someFunc():\n%v\n", justString)

	yetSomeFunc()
	fmt.Printf("result of yetSomeFunc():\n%v\n", justString)

	yetYetSomeFunc()
	fmt.Printf("result of yetYetSomeFunc():\n%v\n", justString)
}

func createHugeString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("[my huge string][моя огромная строка]")
	}
	return sb.String()
}

// возможные последствия - некорректно посчитанное число символов (из-за unicode),
// а также утечка памяти, т.к. в строку срезается только 100 первых байтов.
// остальные - останутся в памяти, т.к. будут иметь на себя ссылки
// для решения можно воспользоваться созданием отдельного среза и копированием переменных туда
