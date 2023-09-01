package main

import "strings"

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }

// Если делаем слайс из массива или другого слайса, то новый слайс будет ссылаться на старый
// до того момента пока не произойдет расширение какого-либо из этих слайсов.
// Следовательно, пока меньший слайс ссылается на больший, больший слайс не будет собран сборщиком
// мусока, даже если он больше не используется.

// В функции someFunc создается очень большая строка, но в переменную justString
// кладется только первые 100 символов этой строки.
// Так как строка представляет слайс байт, то в переменная justString будет ссылаться на исходную строку,
// и исходная строка не будет собрана сборщиком мусока.

// Исправить это можно путем копирования в переменную justString первых 100 символов строки v.
// Таким образом строка v будет собрана после выхода из функции someFunc.

func createHugeString(size int) string {
	return strings.Repeat("x", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 2^10 = 1024

	justString = strings.Clone(v[:100])
}
func main() {
	someFunc()
}
