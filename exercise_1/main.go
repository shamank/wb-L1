package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  string
	sex  string
}

func (h *Human) Say(text string) {
	fmt.Printf("%s said: \"%s\"", h.name, text)
}

type Action struct {
	Human
}

func main() {
	act := Action{}
	act.Say("Hello, world!")
}
