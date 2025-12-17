package cache

import "fmt"

type Lru struct{}

func (l *Lru) Evict(c *Cache) {
	fmt.Println(" [Strategy] Eviciting by LRU")
}
