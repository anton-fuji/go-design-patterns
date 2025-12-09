package main

import (
	"fmt"
	"sync"
)

type DB struct {
	Connection string
}

var (
	instance *DB
	once     sync.Once
)

func GetDBInstance() *DB {
	once.Do(func() {
		fmt.Println("--- Creating DB Instance ---")
		instance = &DB{Connection: "postgresql:5432"}
	})

	return instance
}

func main() {

	var wg sync.WaitGroup

	tasks := 10
	for i := range tasks {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			db := GetDBInstance()
			fmt.Printf("Goroutine: %d \n", id)
			fmt.Printf("DB Addr: %p\n", db)
		}(i)
	}
	wg.Wait()
}
