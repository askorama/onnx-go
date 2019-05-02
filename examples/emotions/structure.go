package main

type emotions []emotion
type emotion struct {
	emotion string
	weight  float32
}

func (e emotions) Len() int           { return len(e) }
func (e emotions) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e emotions) Less(i, j int) bool { return e[i].weight < e[j].weight }
