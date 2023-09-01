package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2....) с использованием конкурентных вычислений.

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// способ 1

func Pow2[T Number](x T) T {
	return x * x
}

func sumOfSquares1[T Number](arr []T, limit int) T {
	var sum T

	ch := make(chan T, limit)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// функция подсчета суммы квадратов, запускается в отдельной горутине
	go func() {
		defer wg.Done()
		for v := range ch {
			sum += v
		}
	}()

	eg := errgroup.Group{}
	eg.SetLimit(limit)

	// запускаем функцию возведения в квадрат, для каждого элемента
	// создаем новую горутину. Одновременно может работать только limit горутин
	for _, v := range arr {
		v := v

		eg.Go(func() error {
			ch <- Pow2(v)
			return nil
		})
	}

	_ = eg.Wait()

	close(ch)

	wg.Wait()
	return sum
}

// способ 2

func gen(nums ...int) <-chan int {
	// создаем канал и пушим в него значения
	in := make(chan int, 128)
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	return in
}

func task(in <-chan int, out chan<- int) {
	//читаем из канала, возводим в квадрат и отправляем в выходной канал
	go func() {
		for v := range in {
			out <- v * v
		}
	}()
}
func calulateSumOfSquares(in <-chan int) int {
	// подсчитываем сумму квадратов
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}

func sumOfSquares2(nums []int, limit int) int {
	in := gen(nums...)

	out := make(chan int, 128)

	wg := &sync.WaitGroup{}

	for i := 1; i <= limit; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			task(in, out)
		}(i)
	}

	// закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(out)
	}()

	return calulateSumOfSquares(out)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Println(sumOfSquares1(arr, 5))

	fmt.Println(sumOfSquares2(arr, 5))

}
