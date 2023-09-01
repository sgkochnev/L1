package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Eat() {
	fmt.Printf("%s eats\n", h.Name)
}

func (h *Human) Sleep() {
	fmt.Printf("%s is asleep\n", h.Name)
}

type Action struct {
	Human
}

func main() {
	tom := Human{
		Name: "Tom",
		Age:  18,
	}

	action := Action{tom}

	action.Eat()
	action.Sleep()

}
