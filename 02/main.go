package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

func Pow2[T Number](x T) T {
	return x * x
}

// Считаем квадрвты чисел из массива. Проблема в том, что при большом количестве чисел
// происходит создание большого количества горутин.
func f0[T Number](arr []T) {
	wg := &sync.WaitGroup{}

	wg.Add(len(arr))

	for _, v := range arr {
		v := v

		go func() {
			defer wg.Done()
			fmt.Println(Pow2(v))
		}()
	}

	wg.Wait()
}

// Добавляем ограничение по количеству одновременно запущенных горутин.
func f1[T Number](arr []T, limit int) {
	wg := &sync.WaitGroup{}

	ch := make(chan struct{}, limit)

	for _, v := range arr {
		ch <- struct{}{}

		wg.Add(1)

		go func(v T) {
			defer wg.Done()
			fmt.Println(Pow2(v))
			<-ch
		}(v)
	}

	close(ch)

	wg.Wait()
}

// С использованием errgroup.
func f2[T Number](arr []T, limit int) {
	eg := errgroup.Group{}

	eg.SetLimit(limit)

	for _, v := range arr {
		v := v

		eg.Go(func() error {
			fmt.Println(Pow2(v))
			return nil
		})
	}

	_ = eg.Wait()
}

// Считаем квадрвты чисел используя воркеры.
func f3(arr []int, limit int) {
	chanVal := make(chan int, 128)

	go func() {
		for _, v := range arr {
			chanVal <- v
		}
		close(chanVal)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(limit)

	for i := 1; i <= limit; i++ {
		go func() {
			defer wg.Done()
			for v := range chanVal {
				fmt.Println(Pow2(v))
			}
		}()
	}

	wg.Wait()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	f0(arr)
	fmt.Println()

	f1(arr, 4)
	fmt.Println()

	f2(arr, 4)
	fmt.Println()

	f3(arr, 2)
}
