package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// partition разбивает массив на две части выбирая опорный элемент,
// все значения меньше опорного перемешаются в левую часть, большие - в правую.
func partition[T any](arr []T, low, high int, less func(i, j int) bool) int {

	pivot := (low + high) / 2

	for l, r := low, high; ; l, r = l+1, r-1 {

		for less(l, pivot) {
			l++
		}

		for less(pivot, r) {
			r--
		}

		if l >= r {
			return r
		}

		// следим за премещением pivot
		if pivot == l {
			pivot = r
		} else if pivot == r {
			pivot = l
		}

		arr[l], arr[r] = arr[r], arr[l]
	}
}

func quicksort[T any](arr []T, less func(i, j int) bool) {
	var qsort func([]T, int, int)

	qsort = func(arr []T, low, high int) {
		if low >= high {
			return
		}

		idx := partition(arr, low, high, less)
		qsort(arr, low, idx)
		qsort(arr, idx+1, high)
	}

	qsort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 9, 8, 2, 5, 4, 6, 15, 8, 1, 3, 7, 0}
	fmt.Println(arr)

	quicksort(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
}
