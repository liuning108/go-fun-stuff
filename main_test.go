package main

import (
	"sync"
	"testing"
)

// BenchmarkCounter_Mutex-12    	   10000	    965368 ns/op	    1045 B/op	      28 allocs/op
func BenchmarkCounter_Mutex(b *testing.B) {
	counter := 0
	mu := sync.Mutex{}

	for i := 0; i < b.N; i++ {
		b.RunParallel(func(pb *testing.PB) {
			n := 0
			for pb.Next() {
				n++
				if n%10000 == 0 {
					//Write
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					//Read
					mu.Lock()
					_ = counter
					mu.Unlock()

				}

			}

		})
	}
}

// BenchmarkCounter_RWMutex-12    	   10000	   1222708 ns/op	    1035 B/op	      28 allocs/op
func BenchmarkCounter_RWMutex(b *testing.B) {
	counter := 0
	mu := sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		b.RunParallel(func(pb *testing.PB) {
			n := 0
			for pb.Next() {
				n++
				if n%10000 == 0 {
					//Write
					mu.Lock()
					counter++
					mu.Unlock()
				} else {
					//Read
					mu.RLock()
					_ = counter
					mu.RUnlock()

				}

			}

		})
	}
}
