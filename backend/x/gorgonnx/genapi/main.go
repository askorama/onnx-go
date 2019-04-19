package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"strings"
	"text/template"
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
		ONNXOpType: "Abs",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Sign",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Ceil",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Floor",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Sin",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Cos",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Exp",
		Arity:      1,
	},
	operation{
		// avoid log as it may conflict with the package
		GorgonnxOp: "logarithm",
		ONNXOpType: "Log",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Log2",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Relu",
		GorgoniaOp: "Rectify",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Neg",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Square",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Sqrt",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Inverse",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Cube",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Tanh",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Sigmoid",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Log1p",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Expm1",
		Arity:      1,
	},
	operation{
		ONNXOpType: "Softplus",
		Arity:      1,
	}}

func main() {
	test := flag.Bool("test", false, "generate test file")
	flag.Parse()
	var t *template.Template
	if *test {
		t = testTmpl
		fmt.Println(testHeader)
	} else {
		t = opTmpl
		fmt.Println(opHeader)
	}
	for _, op := range allOps {
		if op.GorgonnxOp == "" {
			op.GorgonnxOp = strings.ToLower(op.ONNXOpType)
		}
		if op.GorgoniaOp == "" {
			op.GorgoniaOp = op.ONNXOpType
		}

		var buf bytes.Buffer
		if err := t.Execute(&buf, op); err != nil {
			log.Fatal(err)
		}
		p, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal("Cannot format", err)
		}
		fmt.Println(string(p))
	}
}
