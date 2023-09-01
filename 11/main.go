package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](capacity int) Set[V] {
	return make(Set[V], capacity)
}

func NewSetWithValues[V comparable](values ...V) Set[V] {
	set := NewSet[V](len(values))
	for _, v := range values {
		set.Insert(v)
	}
	return set
}

// Положить значение в множество
func (s Set[V]) Insert(value V) {
	s[value] = struct{}{}
}

// Удалить значение из множества
func (s Set[V]) Remove(value V) {
	delete(s, value)
}

// Проверить наличие значения в множестве
func (s Set[V]) Contains(value V) bool {
	_, ok := s[value]
	return ok
}

// Получить количество элементов
func (s Set[V]) Len() int {
	return len(s)
}

// Получить пересечение с другим множеством.
// Возвращает новое множество.
func (s Set[V]) Intersection(set Set[V]) Set[V] {
	intersection := NewSet[V](set.Len())
	for value := range s {
		if ok := set.Contains(value); ok {
			intersection.Insert(value)
		}
	}
	return intersection
}

// Получить разность множеств.
// Возвращает новое множество.
func (s Set[V]) Minus(set *Set[V]) Set[V] {
	minus := NewSet[V](set.Len())
	for value := range s {
		if !set.Contains(value) {
			minus.Insert(value)
		}
	}
	return minus
}

// Получить объединение с другим множеством.
// Возвращает новое множество.
func (s Set[V]) Union(set Set[V]) Set[V] {
	union := NewSet[V](set.Len())
	for value := range s {
		union.Insert(value)
	}
	for value := range set {
		union.Insert(value)
	}
	return union
}

// Получить все значения
func (s Set[V]) Values() []V {
	values := make([]V, 0, s.Len())
	for value := range s {
		values = append(values, value)
	}
	return values
}

func main() {
	set1 := NewSetWithValues(1, 3, 5, 7, 9, 11, 13, 15)
	set2 := NewSetWithValues(1, 3, 6, 10, 15)

	res := set1.Intersection(set2)

	fmt.Println(res.Values())
}
