package main

import (
	"fmt"
	"slices"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseWords(str string) string {
	words := strings.Split(str, " ")

	// l := len(words)
	// for i := 0; i < l/2; i++ {
	// 	j := l - i - 1
	// 	words[i], words[j] = words[j], words[i]
	// }

	slices.Reverse(words) // можно использовать начиная с версии 1.21.0

	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"

	fmt.Println(reverseWords(str))
}
