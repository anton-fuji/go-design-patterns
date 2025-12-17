package main

import "fmt"

type Cache struct {
	storage      map[string]string
	maxCapacity  int
	evictionType string // "lru", "fifo", "lfu" 文字列で判定...
}

// Add はデータを追加するメソッド
// 悪い点: アルゴリズムが増えるたびにこのメソッド内のswitch文が肥大化する
func (c *Cache) Add(key, value string) {
	if len(c.storage) >= c.maxCapacity {
		c.evict()
	}
	c.storage[key] = value
	fmt.Printf("Added: %s\n", key)
}

/* Note
Least Recently Used (LRU)： 使用後最も時間の経った項目を削除
First In, First Out (FIFO)： 最初に作られた項目を削除
Least Frequently Used (LFU)： 利用頻度が最低の項目を削除
*/

// evict は条件分岐でアルゴリズムを切り替える
// 悪い点: 新しいアルゴリズムを追加するたびに、ここを書き換える必要がある
func (c *Cache) evict() {
	switch c.evictionType {
	case "lru":
		fmt.Println("  [Evicting] removing by LRU strategy...")
	case "fifo":
		fmt.Println("  [Evicting] removing by FIFO strategy...")
	case "lfu":
		fmt.Println("  [Evicting] removing by LFU strategy...")
	default:
		fmt.Println("  [Evicting] removing by Default strategy...")
	}

}

func main() {
	c := &Cache{
		storage:      make(map[string]string),
		maxCapacity:  2,
		evictionType: "lru",
	}

	fmt.Println("--- Before Pattern ---")
	c.Add("a", "1")
	c.Add("b", "2")

	// 容量オーバー -> LRUが発動
	c.Add("c", "3")

	// 実行中にアルゴリズムをFIFOに変更（文字列を書き換え）
	c.evictionType = "fifo"
	c.Add("d", "4")
}
