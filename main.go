package main

import (
	"fmt"
)

type Putter interface {
	Put(id int, val any)
}

type Storage interface {
	Get(id int) any
	Put(id int, val any)
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) any {
	return nil
}

func (s *FooStorage) Put(id int, val any) {
	// do nothing
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) any {
	return nil
}

func (s *BarStorage) Put(id int, val any) {
	// do nothing
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) {
	p.Put(id, val)

}

func main() {
	s := &Server{
		store: &FooStorage{},
	}

	updateValue(1, "foo", s.store)

	fmt.Print("Hello, world!")
}
