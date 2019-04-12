package onnx

import (
	"log"

	"gorgonia.org/gorgonia/internal/engine"
	"gorgonia.org/gorgonia/internal/value"
)

// GetOutputValues of the graph
func GetOutputValues(g *Graph) []value.Value {
	output := make([]value.Value, 0)
	it := g.Nodes()
	if it.Next() {
		n := it.Node()
		log.Println(n)
		if g.To(n.ID()).Len() == 0 {
			log.Println("found output")
			output = append(output, n.(*engine.Node).Value())
		}
	}
	return output
}
