package engine

import (
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/node"
	"gorgonia.org/gorgonia/ops"
)

// NewAddOperation ...
func NewAddOperation(leftAxes, rightAxes []byte) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("add: Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		x := children[0]
		y := children[1]

		if leftAxes != nil || rightAxes != nil {
			builder, ok := g.(graph.DirectedWeightedBuilder)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not a DirectedWeightedBuilder")
			}
			_, ok = g.(graph.EdgeRemover)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not an EdgeRemover")
			}

			pattern := newBroadcastPattern(leftAxes, rightAxes)
			broadcastOn := pattern.on()
			switch {
			case len(broadcastOn[0]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), x.ID())
				broadcastedX := builder.NewNode().(*Node)
				broadcastedX.name = n.(*Node).name + "_broadcastedX"
				builder.AddNode(broadcastedX)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedX, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, y, 1.0))

				bcastOp := newBroadcastOperation(second, broadcastOn[0])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedX)
				if err != nil {
					return nil, err
				}
				//x = broadcastedX
			case len(broadcastOn[1]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), y.ID())
				broadcastedY := builder.NewNode().(*Node)
				broadcastedY.name = n.(*Node).name + "_broadcastedY"
				builder.AddNode(broadcastedY)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedY, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, y, 1.0))

				bcastOp := newBroadcastOperation(
					first,
					broadcastOn[1])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedY)
				if err != nil {
					return nil, err
				}
				//y = broadcastedY
			}
		}
		it = getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("AddOperation: Unexpected number of children")
		}
		children = make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(addOpType, children[0], children[1]), nil
	}
}

// HadamardProd perfors a pointwise hadamardprod operation.
func HadamardProd(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(mulOpType, a, b), a, b) }

// NewHadamardProdOperation ...
func NewHadamardProdOperation(leftAxes, rightAxes []byte) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("HadamardProd: Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		x := children[0]
		y := children[1]

		if leftAxes != nil || rightAxes != nil {
			builder, ok := g.(graph.DirectedWeightedBuilder)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not a DirectedWeightedBuilder")
			}
			_, ok = g.(graph.EdgeRemover)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not an EdgeRemover")
			}

			pattern := newBroadcastPattern(leftAxes, rightAxes)
			broadcastOn := pattern.on()
			switch {
			case len(broadcastOn[0]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), x.ID())
				broadcastedX := builder.NewNode().(*Node)
				broadcastedX.name = n.(*Node).name + "_broadcastedX"
				builder.AddNode(broadcastedX)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedX, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, y, 1.0))

				bcastOp := newBroadcastOperation(second, broadcastOn[0])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedX)
				if err != nil {
					return nil, err
				}
				//x = broadcastedX
			case len(broadcastOn[1]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), y.ID())
				broadcastedY := builder.NewNode().(*Node)
				broadcastedY.name = n.(*Node).name + "_broadcastedY"
				builder.AddNode(broadcastedY)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedY, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, y, 1.0))

				bcastOp := newBroadcastOperation(first, broadcastOn[1])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedY)
				if err != nil {
					return nil, err
				}
				//y = broadcastedY
			}
		}
		it = getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("AddOperation: Unexpected number of children")
		}
		children = make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(mulOpType, children[0], children[1]), nil
	}
}
