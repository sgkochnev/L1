package main

import (
	"errors"
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(n int64, i int8, k int8) (int64, error) {

	shift := 0 <= i && i < 64 // проверка на корректность индекса
	bit := k == 0 || k == 1   // проверка на корректность значения

	if !shift || !bit {
		return n, errors.New("can not set bit")
	}

	if k == 0 {
		n = n & ^(1 << i)
	} else {
		n = n | (1 << i)
	}

	return n, nil
}

func main() {

	n := int64(127)

	n, err := setBit(n, 88, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	n, err = setBit(n, 7, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	n, err = setBit(n, 5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
