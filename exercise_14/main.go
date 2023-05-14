package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

func main() {
	var ch interface{} = make(chan interface{})
	arr := []interface{}{55, "23", true, ch}

	for _, val := range arr {
		fmt.Printf("[sc] %v is %s\n", val, simpleCheck(val)) // дефолтное switch-case

		fmt.Printf("[reflect] %v is %s\n", val, reflectCheck(val)) // с помощью пакета reflect

		fmt.Printf("[fmt] %v is %T\n", val, val) // с помощью %T
		fmt.Println("------")
	}
}

func simpleCheck(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}: // однако, если chan определенного типа, то проверка не пройдет..
		return "chan" // и поэтому надо для каждого типа отдельно прописывать
	default:
		return "unknown"
	}
}

func reflectCheck(val interface{}) string {
	return reflect.TypeOf(val).String()
}
