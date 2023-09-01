package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

func worker(ctx context.Context, id int, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d stopped\n", id)
			return
		case v := <-ch:
			fmt.Printf("worker %d, value = %d\n", id, v)
		}
	}
}

func workerPool(ctx context.Context, ch <-chan int, limit int) chan struct{} {

	done := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(limit)

	for i := 1; i <= limit; i++ {
		go func(i int) {
			defer wg.Done()
			worker(ctx, i, ch)
		}(i)
	}

	// Закрываем канал done после завершения всех воркеров.
	go func() {
		wg.Wait()
		close(done)
	}()

	return done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int, 256)

	limit := 16

	done := workerPool(ctx, ch, limit)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- r.Int() % 1000:
			}
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// можно использовать select, так как могут быть другие кейсы завершения программы.
	<-interrupt
	log.Println("Interrupt")

	cancel()
	close(ch)

	// дожидаемся остановки всех воркеров.
	<-done
}
