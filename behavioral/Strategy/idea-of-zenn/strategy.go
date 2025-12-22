package main

type SearchStrategy interface {
	Search(graph map[int][]int, start int, seen []bool) []int
}
