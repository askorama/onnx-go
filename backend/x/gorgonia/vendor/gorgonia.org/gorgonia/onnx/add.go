package onnx

import (
	"errors"

	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/internal/engine"
)

// Add performs a pointwise add.
type Add struct{}

// Constructor to fulfil the interface ...
func (a *Add) Constructor() func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
	return func(g graph.WeightedDirected, n graph.Node) (interface{}, error) {
		it := getOrderedChildren(g, n)
		// Get the shape from the child
		if it.Len() != 2 {
			return nil, errors.New("invalid number of children, expected 2")
		}
		children := make([]*engine.Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*engine.Node)
		}

		x := children[0]
		y := children[1]
		var leftPattern []byte
		var rightPattern []byte
		switch {
		case len(x.Shape()) == 1 && len(y.Shape()) != 1:
			// Need left broadcasting
			// Make an educated guess: find the axis that has the same dimension
			// as x.Shape()[0] and broadcast on all axes of y but this one.
			dims := make([]int, len(x.Shape()))
			for i := 0; i < len(y.Shape()); i++ {
				if y.Shape()[i] != x.Shape()[0] {
					dims[i] = 1
					leftPattern = append(leftPattern, byte(i))
				} else {
					dims[i] = x.Shape()[0]
				}
			}
			if _, ok := g.(graph.DirectedWeightedBuilder); !ok {
				return nil, errors.New("graph is not a builder")
			}
			// Add a reshaped node
			reshaped := g.(graph.DirectedWeightedBuilder).NewNode()
			g.(graph.DirectedWeightedBuilder).AddNode(reshaped)

			w, ok := g.Weight(n.ID(), x.ID())
			if !ok {
				return nil, errors.New("no link found")
			}
			// n -> reshaped
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(n, reshaped, w))
			// reshaped -> x
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(reshaped, x, 0))
			// unlink n -> x
			g.(graph.EdgeRemover).RemoveEdge(n.ID(), x.ID())
			err := g.(*engine.ExprGraph).ApplyOp(engine.Operation(engine.NewReshapeOperation(dims)), reshaped.(*engine.Node))
			if err != nil {
				return nil, err
			}
		case len(y.Shape()) == 1 && len(x.Shape()) != 1:
			// Need right broadcasting
			dims := make([]int, len(x.Shape()))
			for i := 0; i < len(x.Shape()); i++ {
				if x.Shape()[i] != y.Shape()[0] {
					dims[i] = 1
					rightPattern = append(rightPattern, byte(i))
				} else {
					dims[i] = y.Shape()[0]
				}
			}
			var err error
			// Add a reshaped node
			reshaped := g.(graph.DirectedWeightedBuilder).NewNode()
			g.(graph.DirectedWeightedBuilder).AddNode(reshaped)

			w, ok := g.Weight(n.ID(), y.ID())
			if !ok {
				return nil, errors.New("no link found")
			}
			// n -> reshaped
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(n, reshaped, w))
			// reshaped -> x
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(reshaped, y, 0))
			// unlink n -> x
			g.(graph.EdgeRemover).RemoveEdge(n.ID(), y.ID())
			err = g.(*engine.ExprGraph).ApplyOp(engine.Operation(engine.NewReshapeOperation(dims)), reshaped.(*engine.Node))
			if err != nil {
				return nil, err
			}
		case len(y.Shape()) == 3 && len(x.Shape()) == 4:
			// Ugly hack for the mnist model
			dims := make([]int, 4)
			dims[0] = 1
			for i := 0; i < 3; i++ {
				dims[i+1] = y.Shape()[i]
			}
			var err error
			// Add a reshaped node
			reshaped := g.(graph.DirectedWeightedBuilder).NewNode()
			g.(graph.DirectedWeightedBuilder).AddNode(reshaped)

			w, ok := g.Weight(n.ID(), y.ID())
			if !ok {
				return nil, errors.New("no link found")
			}
			// n -> reshaped
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(n, reshaped, w))
			// reshaped -> x
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(reshaped, y, 0))
			// unlink n -> x
			g.(graph.EdgeRemover).RemoveEdge(n.ID(), y.ID())
			err = g.(*engine.ExprGraph).ApplyOp(engine.Operation(engine.NewReshapeOperation(dims)), reshaped.(*engine.Node))
			if err != nil {
				return nil, err
			}
			rightPattern = []byte{0, 2, 3}
		case len(y.Shape()) == 4 && len(x.Shape()) == 3:
			// Ugly hack for the mnist model
			dims := make([]int, 4)
			dims[0] = 1
			for i := 0; i < 3; i++ {
				dims[i+1] = x.Shape()[i]
			}
			var err error
			if _, ok := g.(graph.DirectedWeightedBuilder); !ok {
				return nil, errors.New("graph is not a builder")
			}
			// Add a reshaped node
			reshaped := g.(graph.DirectedWeightedBuilder).NewNode()
			g.(graph.DirectedWeightedBuilder).AddNode(reshaped)

			w, ok := g.Weight(n.ID(), x.ID())
			if !ok {
				return nil, errors.New("no link found")
			}
			// n -> reshaped
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(n, reshaped, w))
			// reshaped -> x
			g.(graph.DirectedWeightedBuilder).SetWeightedEdge(g.(graph.DirectedWeightedBuilder).NewWeightedEdge(reshaped, x, 0))
			// unlink n -> x
			g.(graph.EdgeRemover).RemoveEdge(n.ID(), x.ID())
			err = g.(*engine.ExprGraph).ApplyOp(engine.Operation(engine.NewReshapeOperation(dims)), reshaped.(*engine.Node))
			if err != nil {
				return nil, err
			}

			leftPattern = []byte{0, 2, 3}
		}
		return engine.NewAddOperation(leftPattern, rightPattern)(g, n.(*engine.Node))
	}
}
