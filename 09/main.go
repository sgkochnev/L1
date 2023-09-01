package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

// способ 1

// func gen(in chan int, nums ...int) {
// 	for _, n := range nums {
// 		in <- n
// 	}
// }

// func mul2(in <-chan int, out chan<- int) {
// 	for v := range in {
// 		out <- v * 2
// 	}
// }

// func printX2(out <-chan int) {
// 	for v := range out {
// 		fmt.Println(v)
// 	}
// }

// func conveyor(nums ...int) {
// 	in := make(chan int, 1)
// 	out := make(chan int, 1)

// 	wg := &sync.WaitGroup{}
// 	wg.Add(3)

// 	go func() {
// 		defer close(in)
// 		defer wg.Done()
// 		gen(in, nums...)
// 	}()

// 	go func() {
// 		defer close(out)
// 		defer wg.Done()
// 		mul2(in, out)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		printX2(out)
// 	}()

// 	wg.Wait()
// }

// способ 2
func gen(nums ...int) <-chan int {
	in := make(chan int, 1)
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	return in
}

func mul2(in <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}

func conveyor(nums ...int) {
	in := gen(nums...)
	out := mul2(in)

	for v := range out {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 7, 8, 15, 19, 21, 23, 28, 33}

	conveyor(nums...)

}
