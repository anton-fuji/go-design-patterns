package main

func main() {
	graph := map[int][]int{
		0: []int{1, 2},
		1: []int{3, 4},
		2: []int{5, 6},
	}
	dfsSearcher := NewSearcher("Depth First Search", NewDepthFirstSearchStrategy())
	bfsSearcher := NewSearcher("Breadth First Search", NewBreadthFirstSearchStrategy())

	dfsSearcher.Execute(graph, 0, make([]bool, 7))
	bfsSearcher.Execute(graph, 0, make([]bool, 7))
}
