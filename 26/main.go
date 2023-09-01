package main

import (
	"fmt"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// проверка на уникальность. Сложность O(n) по времени и O(m) по памяти, где
// m — количество различных символов в алфавите
func hasAllUniqueCharacters(str string) bool {
	chars := make(map[rune]struct{})

	for _, char := range str {
		char = unicode.ToLower(char)
		if _, ok := chars[char]; ok {
			return false
		}
		chars[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(hasAllUniqueCharacters("abcd"))      //true
	fmt.Println(hasAllUniqueCharacters("abcdEFGhi")) //true

	fmt.Println(hasAllUniqueCharacters("abCdefAaf")) //false
	fmt.Println(hasAllUniqueCharacters("aabcd"))     //false
	fmt.Println(hasAllUniqueCharacters("aAbcd"))     //false
}
