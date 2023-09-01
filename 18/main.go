package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

// реализация на atomic
type Counter struct {
	value atomic.Int64
}

func NewCounter() *Counter {
	return &Counter{
		value: atomic.Int64{},
	}
}

func (c *Counter) Increment() {
	c.value.Add(1)
}

func (c *Counter) Value() int64 {
	return c.value.Load()
}

// реализация на sync.Mutex менее производительная
// type Counter struct {
// 	mu    sync.Mutex
// 	value int64
// }

// func NewCounter() *Counter {
// 	return &Counter{}
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	c.value++
// 	c.mu.Unlock()
// }

// func (c *Counter) Value() int64 {
// 	return c.value
// }

func main() {
	wg := &sync.WaitGroup{}

	n := 100000

	wg.Add(n)

	c := NewCounter()
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()

	fmt.Println(c.Value())
}
