package main

import "testing"

func BenchmarkWorkerPool_Light_10Workers(b *testing.B) {
	workerSize := 100
	tasks := 10000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(10, 20)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workerSize)
		}
		wp.Shutdown()
	}
}

func BenchmarkWorkerPool_Light_100Workers(b *testing.B) {
	workerSize := 100
	tasks := 10000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(100, 200)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workerSize)
		}
		wp.Shutdown()
	}
}

func BenchmarkWorkerPool_Heavy_10Workers(b *testing.B) {
	workSize := 100000
	tasks := 1000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(10, 20)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workSize)
		}
		wp.Shutdown()
	}
}

func BenchmarkWorkerPool_Heavy_100Workers(b *testing.B) {
	workSize := 100000
	tasks := 1000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(100, 200)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workSize)
		}
		wp.Shutdown()
	}
}

func BenchmarkWorkerPool_ManyTasks_100Workers(b *testing.B) {
	workSize := 1000
	tasks := 10000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(100, 200)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workSize)
		}
		wp.Shutdown()
	}
}

func BenchmarkWorkerPool_ManyTasks_500Workers(b *testing.B) {
	workSize := 1000
	tasks := 10000

	for i := 0; i < b.N; i++ {
		wp := NewWorkerPool(500, 1000)
		wp.Start()
		for j := 0; j < tasks; j++ {
			wp.Submit(workSize)
		}
		wp.Shutdown()
	}
}
