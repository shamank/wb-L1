package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	var (
		foo = 100
		bar = 500
	)

	print(foo, bar)

	// ну почему бы и нет...
	foo, bar = bar, foo
	print(foo, bar)

	// математика
	foo *= bar
	bar = foo / bar
	foo /= bar
	print(foo, bar)

	// с помощью XOR
	foo ^= bar
	bar ^= foo
	foo ^= bar
	print(foo, bar)
}

func print(foo, bar int) {
	fmt.Printf("foo is %v\nbar is %v\n---\n", foo, bar)
}
