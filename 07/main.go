package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

type SyncMap[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func NewSyncMap[K comparable, V any](capacity int) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V, capacity),
	}
}

// Получение ключа из мапы
func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.m[key]
	return v, ok
}

// Положить значение по ключу
func (m *SyncMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

// Удалить значение по ключу
func (m *SyncMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, key)
}

// Получить количество элементов
func (m *SyncMap[K, V]) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}

// Очистить мапу
func (m *SyncMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m = make(map[K]V) // Перезаписываем мапу, старя мапа будет собрана сборщиком мусора.
}

// Получить все ключи
func (m *SyncMap[K, V]) Keys() []K {
	m.mu.Lock()
	defer m.mu.Unlock()
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

// Получить все значения
func (m *SyncMap[K, V]) Values() []V {
	m.mu.Lock()
	defer m.mu.Unlock()
	values := make([]V, 0, len(m.m))
	for _, v := range m.m {
		values = append(values, v)
	}
	return values
}

type Setter interface {
	Set(int64, int)
}

func worker(ctx context.Context, id int, m Setter) {
	r := rand.New(rand.NewSource(time.Now().Unix() + int64(id)))

	d := time.Duration(r.Int()%500 + 100)

	log.Printf("worker %d duration = %d\n", id, d)

	ticker := time.NewTicker(d * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d stopped\n", id)
			return
		case t := <-ticker.C:
			m.Set(t.UnixMilli(), id)
		}
	}
}

// limit - максимальное количество одновременно запущенных горутин
// Возвращает канал сигнализирующий о завершении всех воркеров
func workerPool(ctx context.Context, m Setter, limit int) chan struct{} {
	wg := &sync.WaitGroup{}
	wg.Add(limit)

	for i := 1; i <= limit; i++ {
		go func(i int) {
			defer wg.Done()
			worker(ctx, i, m)
		}(i)
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	return done
}

func writeDataToMap(m Setter) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	done := workerPool(ctx, m, 5)

	<-done
}

func main() {

	m := NewSyncMap[int64, int](0)

	writeDataToMap(m)

	keys := m.Keys()

	for _, k := range keys {
		v, _ := m.Get(k)
		fmt.Printf("%d: %d\n", k, v)
	}

}
