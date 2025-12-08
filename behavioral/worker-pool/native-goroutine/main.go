package main

import (
	"fmt"
	"sync"
)

func doWork(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var wg sync.WaitGroup
	tasks := 1000
	workSize := 1000

	wg.Add(tasks)
	for i := 0; i < tasks; i++ {
		go func(n int) {
			defer wg.Done()
			doWork(n)
		}(workSize)
	}

	wg.Wait()
	fmt.Println("native goroutine done")

}
