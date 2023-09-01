package main

import (
	"cmp"
	"fmt"
)

// Реализовать бинарный поиск встроенными методами языка.

// бинарный поиск реализован в страндартной библиотеке sort.Search и sort.Find
// с версии go 1.21.0 есть еще одана релизация slices.BinarySearch и slices.BinarySearchFunc

// Класический бинарный поиск для cmp.Ordered типов
// Возвращает индекс найденного элемента, если элемент
// не найден вернет -1.
func binarySearch[T cmp.Ordered](arr []T, target T) int {
	l, r := 0, len(arr)
	for l < r {
		m := (r + l) / 2
		switch {
		case arr[m] < target:
			l = m + 1
		case arr[m] > target:
			r = m
		default:
			return m
		}
	}
	return -1
}

// Универсальный бинарный поиск
// Возвращает позицию (индекс) найденного элемента,
// если элемента нет, вернет позицию в которую можно вставить элемент, чтобы сохранить
// слайс в упорядоченном состоянии.
func binSearch(n int, f func(i int) bool) int {
	l, r := 0, n
	for l < r {
		m := (r + l) / 2
		if !f(m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	arr := []int{1, 2, 5, 9, 11, 15, 24, 35}
	target := 36

	idx := binSearch(len(arr), func(i int) bool { return target <= arr[i] })
	if idx < len(arr) && arr[idx] == target {
		fmt.Printf("found: idx = %d, value = %d\n", idx, arr[idx])
	} else {
		fmt.Printf("not found, but %v can be inserted by index %d\n", target, idx)
	}

	idx = binarySearch(arr, target)
	if idx != -1 {
		fmt.Printf("found: idx = %d, value = %d\n", idx, arr[idx])
	} else {
		fmt.Println("not found")
	}
}
