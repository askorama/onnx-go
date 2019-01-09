package main

import (
	"fmt"
	"log"

	onnx "github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/examples/mnist"
	"gorgonia.org/gorgonia"
	"gorgonia.org/gorgonia/debugger/dot"
)

func main() {
	graph := gorgonia.NewGraph()
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
