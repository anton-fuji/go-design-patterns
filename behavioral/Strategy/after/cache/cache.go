package cache

import "fmt"

type EvictionAlgo interface {
	Evict(c *Cache)
}

type Cache struct {
	storage      map[string]string
	maxCapacity  int
	evictionAlgo EvictionAlgo
}

func NewCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		maxCapacity:  2,
		evictionAlgo: e,
	}
}

// 実行時にアルゴリズムを切り替える
func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if len(c.storage) >= c.maxCapacity {
		c.evictionAlgo.Evict(c)
	}
	c.storage[key] = value
	fmt.Printf("Added: %s\n", key)
}
