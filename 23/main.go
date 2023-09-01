package main

import (
	"fmt"

	"slices"
)

// Удалить i-ый элемент из слайса.

// способ 1: без сохранения порядка в слайсе. Работает за O(1).
func removeI1[T any](sl []T, idx int) []T {
	l := len(sl) - 1
	sl[idx] = sl[l]
	return sl[:l]
}

// способ 2: сохраняет порядок в слайсе. Работает за O(n).
func removeI2[T any](sl []T, idx int) []T {
	return append(sl[:idx], sl[idx+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	idx := 3

	sl = removeI1(sl, idx)
	fmt.Println(sl)

	sl = removeI2(sl, idx)
	fmt.Println(sl)

	// можно использовать с версии 1.21.0. Работает за O(n).
	sl = slices.Delete(sl, idx, idx+1)
	fmt.Println(sl)

	// можно использовать с версии 1.21.0. Работает за O(n).
	sl = slices.DeleteFunc(sl, func(v int) bool {
		return v == sl[idx]
	})

	fmt.Println(sl)

}
