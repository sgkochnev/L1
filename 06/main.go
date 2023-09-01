package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// способ 1

func f(ctx context.Context, done <-chan struct{} /*args ...any*/) {
	for {
		select {

		// останока горутины с помощью канала
		case <-done:
			log.Println("f stoped: done")
			return
		// останока горутины с помощью контекста
		case <-ctx.Done():
			log.Println("f stoped: ctx.Done")
			return
		//в других кейса будет логика
		default:
		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		f(ctx, done)
	}()

	wg.Wait()
}
