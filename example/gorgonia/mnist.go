package main

import (
	"fmt"
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gorgonia.org/gorgonia/debugger/dot"
	gorgonnx "gorgonia.org/gorgonia/onnx"
)

func main() {
	graph := gorgonnx.NewGraph()
	err := onnx.Unmarshal(mnist.GetMnist(), graph)
	if err != nil {
		log.Fatal(err)
	}
	b, err := dot.Marshal(graph)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
