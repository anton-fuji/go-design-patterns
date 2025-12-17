package cache

import "fmt"

type Lfu struct{}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println(" [Strategy] Eviciting by LFU")
}
