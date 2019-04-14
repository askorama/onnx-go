package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"strings"
)

var allOps = []operation{
	operation{
		GorgonnxOp:    "hadamardProd",
		ONNXOpType:    "Mul",
		GorgoniaOp:    "HadamardProd",
		Arity:         2,
		Broadcastable: true,
	},
	operation{
		GorgonnxOp:    "hadamardDiv",
		ONNXOpType:    "Div",
		GorgoniaOp:    "HadamardDiv",
		Arity:         2,
		Broadcastable: true,
	},
	operation{
		ONNXOpType:    "Sub",
		Arity:         2,
		Broadcastable: true,
	},
	operation{
		ONNXOpType:    "Add",
		Arity:         2,
		Broadcastable: true,
	},
	operation{
		ONNXOpType: "Cos",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Sin",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Tanh",
		Arity:      1,
	},
}

func main() {
	fmt.Println(`package gorgonnx

import (
	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)`)
	for _, op := range allOps {
		if op.GorgonnxOp == "" {
			op.GorgonnxOp = strings.ToLower(op.ONNXOpType)
		}
		if op.GorgoniaOp == "" {
			op.GorgoniaOp = op.ONNXOpType
		}

		var buf bytes.Buffer
		if err := opTmpl.Execute(&buf, op); err != nil {
			log.Fatal(err)
		}
		p, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal("Cannot format", err)
		}
		fmt.Println(string(p))
	}
}
