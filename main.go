package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *CustomMap[K, V]) Insert(key K, value V) {
	m.data[key] = value
}

func NewCustomMap[k comparable, V any]() *CustomMap[k, V] {
	return &CustomMap[k, V]{
		data: make(map[k]V),
	}
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)
	fmt.Printf("%v\n", m1)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.9)
	m2.Insert(2, 100.99933)
	fmt.Printf("%v\n", m2)

}
