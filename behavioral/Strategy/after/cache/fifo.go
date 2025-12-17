package cache

import "fmt"

type Fifo struct{}

func (l *Fifo) Evict(c *Cache) {
	fmt.Println(" [Strategy] Eviciting by FIFO")
}
