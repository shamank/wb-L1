package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20.

func main() {
	foo := big.NewInt(1234554321)
	bar := big.NewInt(9876556789)
	res := big.NewInt(1)

	fmt.Printf("foo: %v\nbar: %v\n", foo, bar)

	res.Mul(foo, bar)
	fmt.Printf("foo * bar = %v\n", res)

	res.Div(foo, bar)
	fmt.Printf("foo / bar = %v\n", res)

	res.Add(foo, bar)
	fmt.Printf("foo + bar = %v\n", res)

	res.Sub(foo, bar)
	fmt.Printf("foo - bar = %v\n", res)
}
