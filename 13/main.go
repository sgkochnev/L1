package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {

	// способ 1
	// универальный swap
	a, b := 3, 7
	fmt.Printf("1. до 		a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("1. после 	a = %d, b = %d\n", a, b)

	fmt.Println()

	// способ 2
	// работает тоько с числами
	a, b = 3, 7
	fmt.Printf("2. до 		a = %d, b = %d\n", a, b)
	a += b
	b = a - b
	a -= b
	fmt.Printf("2. после 	a = %d, b = %d\n", a, b)
}
