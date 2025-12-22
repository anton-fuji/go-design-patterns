package main

type DepthFirstSearchStrategy struct{}

func NewDepthFirstSearchStrategy() *DepthFirstSearchStrategy {
	return &DepthFirstSearchStrategy{}
}

func (d *DepthFirstSearchStrategy) Search(graph map[int][]int, start int, seen []bool) []int {
	seen[start] = true
	var result []int
	result = append(result, start)

	list := graph[start]
	for _, e := range list {
		if seen[e] {
			continue
		}
		children := d.Search(graph, e, seen)
		result = append(result, children...)
	}

	return result
}
