package main

import (
	"fmt"
	"slices"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverse(str string) string {
	rev := []rune(str)

	// l := len(rev)
	// for i := 0; i < l/2; i++ {
	// 	j := l - i - 1
	// 	rev[i], rev[j] = rev[j], rev[i]
	// }

	slices.Reverse(rev) // можно использовать начиная с версии 1.21.0

	return string(rev)
}

func main() {
	str := "главрыба"

	fmt.Println(reverse(str))
}
