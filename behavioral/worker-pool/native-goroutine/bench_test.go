package main

import (
	"sync"
	"testing"
)

func BenchmarkNativeGoroutine_Light(b *testing.B) {
	workSize := 100
	tasks := 1000

	for b.Loop() {
		var wg sync.WaitGroup
		wg.Add(tasks)
		for j := 0; j < tasks; j++ {
			go func(n int) {
				defer wg.Done()
				doWork(n)
			}(workSize)
		}
		wg.Wait()
	}
}

func BenchmarkNativeGoroutine_Medium(b *testing.B) {
	workSize := 10000
	tasks := 1000

	for b.Loop() {
		var wg sync.WaitGroup
		wg.Add(tasks)
		for j := 0; j < tasks; j++ {
			go func(n int) {
				defer wg.Done()
				doWork(n)
			}(workSize)
		}
		wg.Wait()
	}
}

func BenchmarkNativeGoroutine_Heavy(b *testing.B) {
	workSize := 1000000
	tasks := 1000

	for b.Loop() {
		var wg sync.WaitGroup
		wg.Add(tasks)
		for j := 0; j < tasks; j++ {
			go func(n int) {
				defer wg.Done()
				doWork(n)
			}(workSize)
		}
		wg.Wait()
	}
}
