package main

import (
	"fmt"
	"sync"
)

type Job struct {
	n int
}

type WorkerPool struct {
	workerNum int
	jobQueue  chan Job
	wg        sync.WaitGroup
	once      sync.Once
}

func NewWorkerPool(workerNum, bufSize int) *WorkerPool {
	return &WorkerPool{
		workerNum: workerNum,
		jobQueue:  make(chan Job, bufSize),
	}
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done() // worker終了時にカウントダウン
	for job := range wp.jobQueue {
		doWork(job.n)
	}
}

func (wp *WorkerPool) Start() {
	wp.once.Do(func() {
		wp.wg.Add(wp.workerNum) // workerの数だけAdd
		for i := 0; i < wp.workerNum; i++ {
			go wp.worker()
		}
	})
}

func (wp *WorkerPool) Submit(n int) {
	wp.jobQueue <- Job{n: n}
}

// 新しいJobの受付を停止し、全てのworkerが終了するまで待つ
func (wp *WorkerPool) Shutdown() {
	close(wp.jobQueue)
	wp.wg.Wait()
}

func doWork(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	tasks := 10000
	workSize := 10000
	workerNum := 100

	wp := NewWorkerPool(workerNum, workerNum*2)
	wp.Start()

	for i := 0; i < tasks; i++ {
		wp.Submit(workSize)
	}

	wp.Shutdown()
	fmt.Println("worke pool done")
}
