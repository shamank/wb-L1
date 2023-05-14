package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type Client struct {
}

// клиент может использовать только те языки, которые знает (которые реализуют интерфейс Developer)
func (c *Client) Work(dev Developer) {

	dev.PrintHelloWorld()
}

type Developer interface {
	PrintHelloWorld()
}

type PythonDeveloper struct {
}

func (p *PythonDeveloper) PrintHelloWorld() {
	fmt.Printf("print('Hello, world!)\n'")
}

type AssemblerDeveloper struct {
}

func (a *AssemblerDeveloper) HardWork() {
	fmt.Printf(`
				.model tiny
				.code
				ORG 100h
					
				begin:	
				  MOV  AH,  9          ; Функция вывода строки на экран 
				  MOV  DX,  OFFSET Msg ; Адрес строки
				  INT  21h             ; Выполнить функцию
				
				  RET                  ; Вернуться в операционную систему
						
				Msg DB 'Hello, World!!!$'   ; Строка для вывода
				
				END begin
	`)
}

type AssemblerAdapter struct {
	aDeveloper *AssemblerDeveloper
}

func (a *AssemblerAdapter) PrintHelloWorld() {
	a.aDeveloper.HardWork()
}

func main() {

	// клиент знает пайтона, но с ассемблером у него проблемы
	client := &Client{}

	pythonDev := &PythonDeveloper{}

	// он может спокойно писать с помощью пайтона
	client.Work(pythonDev)

	// но чтобы написать программу на ассемблере, ему нужен *адаптер*
	assemblerDev := &AssemblerDeveloper{}
	assemblerAdapter := &AssemblerAdapter{assemblerDev}

	// так наш клиент научился использовать ассемблер!!!
	client.Work(assemblerAdapter)
}
