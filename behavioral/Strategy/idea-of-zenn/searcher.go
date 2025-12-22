package main

import (
	"encoding/json"
	"fmt"
)

type Searcher struct {
	searcherName string
	strategy     SearchStrategy
}

func NewSearcher(searcherName string, s SearchStrategy) *Searcher {
	return &Searcher{searcherName: searcherName, strategy: s}
}

func (s *Searcher) Execute(graph Graph, start int, seen []bool) {
	res := s.strategy.Search(graph, start, seen)
	str, _ := json.Marshal(res)
	fmt.Printf("name: %s, result: %s\n", s.searcherName, string(str))
}
