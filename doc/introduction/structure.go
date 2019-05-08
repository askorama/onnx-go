package main

type results []result
type result struct {
	result string
	weight float32
}

func (e results) Len() int           { return len(e) }
func (e results) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e results) Less(i, j int) bool { return e[i].weight < e[j].weight }
