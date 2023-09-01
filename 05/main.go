package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

const N = 3

func writeInChan(ctx context.Context, ch chan<- int) {
	val := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("writeInChan stopped")
			return
		case ch <- val:
			val++
		}
	}
}

func readFromChan(ctx context.Context, ch <-chan int) {
	var v int
	for {
		select {
		case <-ctx.Done():
			log.Println("readFromChan stopped")
			log.Printf("value = %d", v)
			return
		case v = <-ch:
			// здесь какая-то логика
			// log.Println("value = ", v)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*N)
	defer cancel()

	ch := make(chan int, 64)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		writeInChan(ctx, ch)
	}()

	go func() {
		defer wg.Done()
		readFromChan(ctx, ch)
	}()

	wg.Wait()
	close(ch)
}
