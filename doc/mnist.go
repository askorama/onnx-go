package main

import (
	"fmt"
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gonum.org/v1/gonum/graph/encoding/dot"
)

func main() {
	graph := newSimpleGraph()
	err := onnx.Unmarshal(mnist.GetMnist(), graph)
	if err != nil {
		log.Fatal(err)
	}
	b, err := dot.Marshal(graph, "name", "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
