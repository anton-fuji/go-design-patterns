package main

type BreadthFirstSearchStrategy struct {
	queue []int
}

func NewBreadthFirstSearchStrategy() *BreadthFirstSearchStrategy {
	return &BreadthFirstSearchStrategy{queue: make([]int, 0)}
}

func (b *BreadthFirstSearchStrategy) Search(graph map[int][]int, start int, seen []bool) []int {
	seen[start] = true
	b.queue = append(b.queue, start)

	var result []int
	for len(b.queue) != 0 {
		v := b.queue[0]
		b.queue = b.queue[1:]

		result = append(result, v)
		list := graph[v]
		for _, e := range list {
			if seen[e] {
				continue
			}
			b.queue = append(b.queue, e)
		}
	}
	return result
}
