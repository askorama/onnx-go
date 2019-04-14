package main

import "text/template"

var opTmpl = template.Must(template.New("op").Funcs(iterator).Parse(opTemplate))

type operation struct {
	GorgonnxOp    string
	ONNXOpType    string
	GorgoniaOp    string
	Arity         int
	Broadcastable bool
}

const opTemplate = `
type {{ .GorgonnxOp }} struct{}

func init() {
	register("{{ .ONNXOpType }}", &{{ .GorgonnxOp }}{})
}

func (a *{{ .GorgonnxOp }}) apply(g *Graph, n *Node) error {
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, {{ .Arity }})
	if err != nil {
		return err
	}
	{{ if .Broadcastable }}
	if len(children[0].gorgoniaNode.Shape()) != len(children[1].gorgoniaNode.Shape()) {
		return &onnx.ErrNotImplemented{
			Operator: n.operation.Name,
			Message:  "broadcast not yet implemented",
		}

	}
	{{ end }}
	n.gorgoniaNode, err = gorgonia.{{ .GorgoniaOp }}(
		{{- range $val := Iterate .Arity }}
		children[{{ $val }}].gorgoniaNode, 
		{{- end }}
	)
	return err
}

func (a *{{ .GorgonnxOp }}) init(o onnx.Operation) error {
	return nil
}
`

var iterator = template.FuncMap{
	"Iterate": func(count int) []int {
		var i int
		var Items []int
		for i = 0; i < count; i++ {
			Items = append(Items, i)
		}
		return Items
	},
}
