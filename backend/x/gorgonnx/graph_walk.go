package gorgonnx

import (
	"errors"
	"fmt"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

// populateExprgraph by walking through the graph
func (g *Graph) populateExprgraph() error {
	if len(g.groups) == 0 {
		return errors.New("cannot populate the graph because ApplyOperation have not been called")
	}

	// Walk the graph
	itN := g.Nodes()
	for itN.Next() {
		// if the node is a "tensor", set it!
		n := itN.Node().(*Node)
		if n.gorgoniaNode == nil && n.operation == nil {
			n.gorgoniaNode = &gorgonia.Node{}
			if n.t != nil {
				n.gorgoniaNode = gorgonia.NodeFromAny(g.exprgraph, n.t, gorgonia.WithName(getUniqNodeName("node")))
			}
		}
	}
	nodes := make([][]*Node, len(g.groups))
	copy(nodes, g.groups)
	for len(nodes) > 0 {
		initialLen := len(nodes)
		for i := 0; i < len(nodes); i++ {
			nilChild := false
			for _, n := range nodes[i] {
				//if n.operation != nil {
				children := getOrderedChildren(g.g, n)
				for j := 0; j < len(children); j++ {
					if children[j].gorgoniaNode == nil {
						nilChild = true
						break
					}
				}
				//}
			}
			if nilChild {
				continue
			}
			err := g.applyOperation(nodes[i]...)
			if err != nil {
				return err
			}
			nodes = append(nodes[:i], nodes[i+1:]...)
		}
		if len(nodes) == initialLen {
			return errors.New("infinite loop")
		}
	}
	return nil
}

// applyOperation creates a new node on the exprgraph
func (g *Graph) applyOperation(n ...*Node) error {
	// Is this node already in the ExprGraph?
	if n[0].gorgoniaNode != nil {
		return fmt.Errorf("unsupported case: node is already in the exprgraph")
	}
	var op operator
	var opC func() operator
	var ok bool
	if opC, ok = operators[n[0].operation.Name]; !ok {
		return &onnx.ErrNotImplemented{
			Operator: n[0].operation.Name,
		}
	}
	op = opC()
	err := op.init(*n[0].operation)
	if err != nil {
		return err
	}
	return op.apply(g, n...)
}
